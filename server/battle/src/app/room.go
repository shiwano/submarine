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
	id            int64
	webAPI        *webAPI.WebAPI
	info          *battleAPI.Room
	sessions      map[int64]*Session
	battle        *battle.Battle
	closeHandler  func(*Room)
	startBattleCh chan *Session
	joinCh        chan *Session
	leaveCh       chan *Session
	closeCh       chan struct{}
	isClosed      *abool.AtomicBool
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
		battle:        battle.New(time.Second*60, stageMesh),
		startBattleCh: make(chan *Session, 1),
		joinCh:        make(chan *Session, 4),
		leaveCh:       make(chan *Session, 4),
		closeCh:       make(chan struct{}, 1),
		isClosed:      abool.New(),
	}

	go room.run()
	return room, nil
}

func (r *Room) run() {
	logger.Log.Infof("Room(%v) opened", r.id)

loop:
	for {
		select {
		case session := <-r.startBattleCh:
			r.startBattle(session)
		case session := <-r.joinCh:
			r.join(session)
			r.broadcastRoom()
			session.synchronizeTime()
		case session := <-r.leaveCh:
			r.leave(session)
			r.broadcastRoom()
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
	return &api.Room{Id: r.id, Members: members}
}

func (r *Room) broadcastRoom() {
	typhenType := r.toRoomAPIType()
	for _, s := range r.sessions {
		s.api.Battle.SendRoom(typhenType)
	}
}

func (r *Room) startBattle(session *Session) {
	// TODO: Validate that can the session starts the battle.
	if r.battle.Start() {
		logger.Log.Infof("Room(%v)'s battle started", r.id)
	}
}

func (r *Room) join(session *Session) {
	logger.Log.Infof("Session(%v) joined into Room(%v)", session.id, r.id)
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.leaveCh <- session
	}
	r.battle.EnterUser(session.id)
}

func (r *Room) leave(session *Session) {
	logger.Log.Infof("Session(%v) leaved from Room(%v)", session.id, r.id)
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
	r.battle.LeaveUser(session.id)
}

func (r *Room) close() {
	for c := 1; true; c++ {
		if _, err := r.webAPI.Battle.CloseRoom(r.id); err != nil {
			logger.Log.Errorf("Room(%v) failed %v times to use closeRoom API: %v", r.id, c, err)
			time.Sleep(time.Duration(c) * time.Second)
			continue
		}
		break
	}
	logger.Log.Infof("Room(%v) closed", r.id)
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
