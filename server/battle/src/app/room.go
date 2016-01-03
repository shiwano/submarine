package main

import (
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	webapi "app/typhenapi/web/submarine"
	"errors"
)

// Room represents a network group for battle.
type Room struct {
	id           int64
	webAPI       *webapi.WebAPI
	info         *battle.Room
	sessions     map[int64]*Session
	closeHandler func(*Room)
	join         chan *Session
	leave        chan *Session
	close        chan struct{}
}

func newRoom(id int64) (*Room, error) {
	webAPI := NewWebAPI("http://localhost:3000")

	// TODO: Validation for creatable the room in the battle server.
	res, err := webAPI.Battle.FindRoom(id)
	if err != nil {
		return nil, err
	}
	if res.Room == nil {
		return nil, errors.New("No room found.")
	}

	room := &Room{
		id:       id,
		webAPI:   webAPI,
		info:     res.Room,
		sessions: make(map[int64]*Session),
		join:     make(chan *Session, 4),
		leave:    make(chan *Session, 4),
		close:    make(chan struct{}),
	}

	go room.run()
	return room, nil
}

func (r *Room) run() {
loop:
	for {
		select {
		case session := <-r.join:
			r._join(session)
		case session := <-r.leave:
			r._leave(session)
		case <-r.close:
			r._close()
			break loop
		}
	}

	if r.closeHandler != nil {
		r.closeHandler(r)
	}
}

func (r *Room) findRoomMember(roomKey string) *battle.RoomMember {
	for _, m := range r.info.Members {
		if m.RoomKey == roomKey {
			return m
		}
	}
	return nil
}

func (r *Room) toRoomAPIType() *submarine.Room {
	members := make([]*submarine.User, len(r.sessions))
	for i, s := range r.sessions {
		members[i] = s.toUserAPIType()
	}
	return &submarine.Room{Id: r.id, Members: members}
}

func (r *Room) _join(session *Session) {
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.leave <- session
	}
}

func (r *Room) _leave(session *Session) {
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
}

func (r *Room) _close() {
	for _, session := range r.sessions {
		r._leave(session)
		session.close()
	}
}
