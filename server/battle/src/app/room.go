package main

import (
	"app/typhen_api/core"
	"app/typhen_api/type/submarine/battle"
	"github.com/olahol/melody"
)

// Room represents a network group for battle.
type Room struct {
	serializer *typhenapi.Serializer
	sessions   map[*melody.Session]*Session
}

func (room *Room) join(rawSession *melody.Session) {
	session := newSession(rawSession, room.serializer)

	session.Dispatcher.Battle.Ping = func(ping *battle.Ping) {
		session.send(ping)
	}

	room.sessions[rawSession] = session
}

func (room *Room) leave(rawSession *melody.Session) {
	delete(room.sessions, rawSession)
}

func (room *Room) handleMessage(rawSession *melody.Session, data []byte) {
	room.sessions[rawSession].handleMessage(data)
}

func (room *Room) close() {
	for _, session := range room.sessions {
		session.Close()
	}
	room.sessions = nil
}
