package server

import (
	"fmt"
	"time"

	"github.com/tevino/abool"

	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	webAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	"github.com/shiwano/submarine/server/battle/server/battle"
	"github.com/shiwano/submarine/server/battle/server/config"
	"github.com/shiwano/submarine/server/battle/server/logger"
	"github.com/shiwano/submarine/server/battle/server/resource"
)

type room struct {
	id               int64
	webAPI           *webAPI.WebAPI
	info             *battleAPI.Room
	sessions         map[int64]*session
	bots             map[int64]*api.Bot
	battle           *battle.Battle
	closeHandler     func(*room)
	isClosed         *abool.AtomicBool
	lastCreatedBotID int64
	startBattleCh    chan *session
	addBotCh         chan struct{}
	removeBotCh      chan int64
	joinCh           chan *session
	leaveCh          chan *session
	closeCh          chan struct{}
}

func newRoom(id int64) (*room, error) {
	webAPI := newWebAPI(config.Config.ApiServerBaseUri)

	// TODO: Validate whether the battle server can create the room.
	res, err := webAPI.Battle.FindRoom(id)
	if err != nil {
		return nil, err
	}
	if res.Room == nil {
		return nil, fmt.Errorf("No room(%v) found", id)
	}

	// TODO: Load the specified stage mesh and the light map.
	stageMesh, err := resource.Loader.LoadMesh(1)
	if err != nil {
		return nil, err
	}
	lightMap, err := resource.Loader.LoadLightMap(1)
	if err != nil {
		return nil, err
	}

	r := &room{
		id:            id,
		webAPI:        webAPI,
		info:          res.Room,
		sessions:      make(map[int64]*session),
		bots:          make(map[int64]*api.Bot),
		battle:        battle.New(time.Second*300, stageMesh, lightMap),
		startBattleCh: make(chan *session, 1),
		addBotCh:      make(chan struct{}, 1),
		removeBotCh:   make(chan int64, 1),
		joinCh:        make(chan *session, 4),
		leaveCh:       make(chan *session, 4),
		closeCh:       make(chan struct{}, 1),
		isClosed:      abool.New(),
	}

	go r.run()
	return r, nil
}

func (r *room) String() string {
	return fmt.Sprintf("Room(%v)", r.id)
}

func (r *room) run() {
	logger.Log.Infof("%v opened", r)

loop:
	for {
		select {
		case session := <-r.startBattleCh:
			r.startBattle(session)
		case <-r.addBotCh:
			r.addBot()
		case botID := <-r.removeBotCh:
			r.removeBot(botID)
		case session := <-r.joinCh:
			r.join(session)
		case session := <-r.leaveCh:
			r.leave(session)
		case <-r.closeCh:
			r.close()
			break loop
		case output := <-r.battle.Gateway.Output:
			r.onBattleOutputReceive(output)
		}
	}

	r.isClosed.SetTo(true)

	if r.closeHandler != nil {
		r.closeHandler(r)
	}
}

func (r *room) toRoomAPIType() *api.Room {
	members := make([]*api.User, len(r.sessions))
	i := 0
	for _, s := range r.sessions {
		members[i] = s.toUserAPIType()
		i++
	}
	bots := make([]*api.Bot, len(r.bots))
	i = 0
	for _, b := range r.bots {
		bots[i] = b
		i++
	}
	return &api.Room{Id: r.id, Members: members, Bots: bots}
}

func (r *room) broadcastRoom() {
	message := r.toRoomAPIType()
	for _, s := range r.sessions {
		s.api.Battle.SendRoom(message)
	}
}

func (r *room) startBattle(s *session) {
	// TODO: Validate that can the session starts the battle.
	if r.battle.Start() {
		logger.Log.Infof("%v's battle started", r)
	}
}

func (r *room) addBot() {
	r.lastCreatedBotID--
	bot := &api.Bot{Id: r.lastCreatedBotID, Name: "BOT"}
	if r.battle.EnterBot(bot) {
		r.bots[bot.Id] = bot
		r.broadcastRoom()
		logger.Log.Infof("Bot(%v) is added to %v", bot.Id, r)
	}
}

func (r *room) removeBot(botID int64) {
	if bot, ok := r.bots[botID]; ok {
		if r.battle.LeaveBot(bot) {
			delete(r.bots, bot.Id)
			r.broadcastRoom()
			logger.Log.Infof("Bot(%v) is removed from %v", bot.Id, r)
		}
	}
}

func (r *room) join(s *session) {
	logger.Log.Infof("%v joined into %v", s, r)
	r.sessions[s.id] = s
	s.room = r
	s.disconnectHandler = func(s *session) {
		r.leaveCh <- s
	}
	r.battle.EnterUser(s.id)
	s.synchronizeTime()
	r.broadcastRoom()
}

func (r *room) leave(s *session) {
	logger.Log.Infof("%v leaved from %v", s, r)
	s.disconnectHandler = nil
	s.room = nil
	delete(r.sessions, s.id)
	r.battle.LeaveUser(s.id)
	r.broadcastRoom()
}

func (r *room) close() {
	for c := 1; true; c++ {
		if _, err := r.webAPI.Battle.CloseRoom(r.id); err != nil {
			logger.Log.Errorf("%v failed %v times to use closeRoom API: %v", r, c, err)
			time.Sleep(time.Duration(c) * time.Second)
			continue
		}
		break
	}
	logger.Log.Infof("%v closed", r)
	r.battle.Close()
	for _, session := range r.sessions {
		r.leave(session)
		session.close()
	}
}

func (r *room) sendBattleInput(userID int64, message typhenapi.Type) {
	r.battle.Gateway.InputMessage(userID, message)
}

func (r *room) onBattleOutputReceive(output *battle.GatewayOutput) {
	if output.UserIDs == nil {
		for _, s := range r.sessions {
			s.api.Battle.Send(output.Message)
		}
	} else {
		for _, s := range r.sessions {
			for _, userID := range output.UserIDs {
				if s.id == userID {
					s.api.Battle.Send(output.Message)
				}
			}
		}
	}
	if output.IsFinishMessage {
		go func() { r.closeCh <- struct{}{} }()
	}
}
