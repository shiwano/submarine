package main

import (
	battleLogic "app/battle"
	"app/logger"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	webapi "app/typhenapi/web/submarine"
	websocketapi "app/typhenapi/websocket/submarine/battle"
	"fmt"
	"time"
)

// Room represents a network group for battle.
type Room struct {
	id           int64
	webAPI       *webapi.WebAPI
	info         *battle.Room
	sessions     map[int64]*Session
	battle       *battleLogic.Battle
	closeHandler func(*Room)
	join         chan *Session
	leave        chan *Session
	close        chan struct{}
	isClosed     bool
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

	room := &Room{
		id:       id,
		webAPI:   webAPI,
		info:     res.Room,
		sessions: make(map[int64]*Session),
		battle:   battleLogic.New(time.Second * 60),
		join:     make(chan *Session, 4),
		leave:    make(chan *Session, 4),
		close:    make(chan struct{}, 1),
	}

	go room.run()
	return room, nil
}

func (r *Room) run() {
	logger.Log.Infof("Room(%v) opened", r.id)

loop:
	for {
		select {
		case session := <-r.join:
			r._join(session)
			r.broadcastRoom()
			session.synchronizeTime()
		case session := <-r.leave:
			r._leave(session)
			r.broadcastRoom()
		case <-r.close:
			r._close()
			break loop
		case output := <-r.battle.Gateway.Output:
			r.onBattleOutputReceive(output)
		}
	}

	r.isClosed = true
	close(r.join)
	close(r.leave)
	close(r.close)

	if r.closeHandler != nil {
		r.closeHandler(r)
	}
}

func (r *Room) toRoomAPIType() *submarine.Room {
	members := make([]*submarine.User, len(r.sessions))
	i := 0
	for _, s := range r.sessions {
		members[i] = s.toUserAPIType()
		i++
	}
	return &submarine.Room{Id: r.id, Members: members}
}

func (r *Room) broadcastRoom() {
	typhenType := r.toRoomAPIType()
	for _, s := range r.sessions {
		s.api.Battle.SendRoom(typhenType)
	}
}

func (r *Room) _join(session *Session) {
	logger.Log.Infof("Session(%v) joined into Room(%v)", session.id, r.id)
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.leave <- session
	}
	if !r.battle.IsStarted {
		r.battle.CreateSubmarineUnlessExists(session.id)
	}

	// TODO: Add relevant room members counting.
	if len(r.sessions) >= 1 {
		r.battle.Start()
	}
}

func (r *Room) _leave(session *Session) {
	logger.Log.Infof("Session(%v) leaved from Room(%v)", session.id, r.id)
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
}

func (r *Room) _close() {
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
		r._leave(session)
		session.close()
	}
}

func (r *Room) sendBattleInput(userID int64, message typhenapi.Type) {
	r.battle.Gateway.InputMessage(userID, message)
}

func (r *Room) onBattleOutputReceive(output *battleLogic.GatewayOutput) {
	var message *typhenapi.Message
	if output.UserIDs == nil {
		for _, s := range r.sessions {
			message, _ = s.api.Battle.Send(output.Message)
		}
	} else {
		for _, s := range r.sessions {
			for _, userID := range output.UserIDs {
				if s.id == userID {
					message, _ = s.api.Battle.Send(output.Message)
				}
			}
		}
	}
	if message.Type == websocketapi.MessageType_Finish {
		go func() { r.close <- struct{}{} }()
	}
}
