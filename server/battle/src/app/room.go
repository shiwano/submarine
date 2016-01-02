package main

import (
	webapi "app/typhenapi/web/submarine"
)

// Room represents a network group for battle.
type Room struct {
	id           uint64
	webAPI       *webapi.WebAPI
	sessions     map[uint64]*Session
	closeHandler func(*Room)
	Join         chan *Session
	Leave        chan *Session
	Close        chan struct{}
}

func newRoom(id uint64) *Room {
	room := &Room{
		id:       id,
		webAPI:   NewWebAPI("http://localhost:3000"),
		sessions: make(map[uint64]*Session),
		Join:     make(chan *Session),
		Leave:    make(chan *Session),
		Close:    make(chan struct{}),
	}
	go room.run()
	return room
}

func (r *Room) run() {
loop:
	for {
		select {
		case session := <-r.Join:
			r.join(session)
		case session := <-r.Leave:
			r.leave(session)
		case <-r.Close:
			r.leaveAndCloseSessions()
			break loop
		}
	}

	if r.closeHandler != nil {
		r.closeHandler(r)
	}
}

func (r *Room) join(session *Session) {
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.Leave <- session
	}
}

func (r *Room) leave(session *Session) {
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
}

func (r *Room) leaveAndCloseSessions() {
	for _, session := range r.sessions {
		r.leave(session)
		session.close()
	}
}
