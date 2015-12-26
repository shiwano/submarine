package main

import (
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	api "app/typhenapi/websocket/submarine"
	"github.com/olahol/melody"
)

// Session represents a network session that has user infos.
type Session struct {
	*melody.Session
	api *api.WebSocketAPI
}

// Send sends raw message data.
func (session *Session) Send(msg []byte) {
	session.WriteBinary(msg)
}

func newSession(melodySession *melody.Session, serializer *typhenapi.Serializer) *Session {
	session := &Session{melodySession, nil}
	session.api = api.New(session, serializer, func(m []byte, err error) { Log.Error(err) })
	session.api.Battle.Ping = session.onPingReceive
	return session
}

func (session *Session) onPingReceive(message *battle.Ping) {
	Log.Info("ping")
	session.api.Battle.SendPing(message)
}

func (session *Session) handleMessage(data []byte) {
	session.api.DispatchMessageEvent(data)
}
