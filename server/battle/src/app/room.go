package main

import (
	"app/battle"
	"app/logger"
	"app/resource"
	"app/typhenapi/core"
	api "app/typhenapi/type/submarine"
	battleAPI "app/typhenapi/type/submarine/battle"
	webAPI "app/typhenapi/web/submarine"
	"fmt"
	"github.com/tevino/abool"
	"time"
)

// Room represents a network group for battle.
type Room struct {
	id               int64
	webAPI           *webAPI.WebAPI
	info             *battleAPI.Room
	sessions         map[int64]*Session
	bots             map[int64]*api.Bot
	battle           *battle.Battle
	closeHandler     func(*Room)
	isClosed         *abool.AtomicBool
	lastCreatedBotID int64
	startBattleCh    chan *Session
	addBotCh         chan struct{}
	removeBotCh      chan int64
	joinCh           chan *Session
	leaveCh          chan *Session
	closeCh          chan struct{}
}

func newRoom(id int64) (*Room, error) {
	webAPI := NewWebAPI("http://localhost:3000")

	// TODO: Validation for creatable the room in the battle server.
	res, err := webAPI.Battle.FindRoom(id)
	if err != nil {
		return nil, err
	}
	if res.Room == nil {
		return nil, fmt.Errorf("No room(%v) found", id)
	}

	// TODO: Specify relevant stage code.
	stageMesh, err := resource.Loader.LoadStageMesh(1)
	if err != nil {
		return nil, err
	}

	room := &Room{
		id:            id,
		webAPI:        webAPI,
		info:          res.Room,
		sessions:      make(map[int64]*Session),
		bots:          make(map[int64]*api.Bot),
		battle:        battle.New(time.Second*300, stageMesh),
		startBattleCh: make(chan *Session, 1),
		addBotCh:      make(chan struct{}, 1),
		removeBotCh:   make(chan int64, 1),
		joinCh:        make(chan *Session, 4),
		leaveCh:       make(chan *Session, 4),
		closeCh:       make(chan struct{}, 1),
		isClosed:      abool.New(),
	}

	go room.run()
	return room, nil
}

func (r *Room) String() string {
	return fmt.Sprintf("Room(%v)", r.id)
}

func (r *Room) run() {
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

func (r *Room) toRoomAPIType() *api.Room {
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

func (r *Room) broadcastRoom() {
	message := r.toRoomAPIType()
	for _, s := range r.sessions {
		s.api.Battle.SendRoom(message)
	}
}

func (r *Room) startBattle(session *Session) {
	// TODO: Validate that can the session starts the battle.
	if r.battle.Start() {
		logger.Log.Infof("%v's battle started", r)
	}
}

func (r *Room) addBot() {
	r.lastCreatedBotID--
	bot := &api.Bot{Id: r.lastCreatedBotID, Name: "BOT"}
	if r.battle.EnterBot(bot) {
		r.bots[bot.Id] = bot
		r.broadcastRoom()
		logger.Log.Infof("Bot(%v) is added to %v", bot.Id, r)
	}
}

func (r *Room) removeBot(botID int64) {
	if bot, ok := r.bots[botID]; ok {
		if r.battle.LeaveBot(bot) {
			delete(r.bots, bot.Id)
			r.broadcastRoom()
			logger.Log.Infof("Bot(%v) is removed from %v", bot.Id, r)
		}
	}
}

func (r *Room) join(session *Session) {
	logger.Log.Infof("%v joined into %v", session, r)
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.leaveCh <- session
	}
	r.battle.EnterUser(session.id)
	session.synchronizeTime()
	r.broadcastRoom()
}

func (r *Room) leave(session *Session) {
	logger.Log.Infof("%v leaved from %v", session, r)
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
	r.battle.LeaveUser(session.id)
	r.broadcastRoom()
}

func (r *Room) close() {
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

func (r *Room) sendBattleInput(userID int64, message typhenapi.Type) {
	r.battle.Gateway.InputMessage(userID, message)
}

func (r *Room) onBattleOutputReceive(output *battle.GatewayOutput) {
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
