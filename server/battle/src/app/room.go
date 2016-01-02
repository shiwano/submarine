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
	join         chan *Session
	leave        chan *Session
	close        chan struct{}
}

func newRoom(id uint64) *Room {
	room := &Room{
		id:       id,
		webAPI:   NewWebAPI("http://localhost:3000"),
		sessions: make(map[uint64]*Session),
		join:     make(chan *Session, 4),
		leave:    make(chan *Session, 4),
		close:    make(chan struct{}),
	}
	go room.run()
	return room
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
