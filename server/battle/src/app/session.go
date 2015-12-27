package main

import (
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	api "app/typhenapi/websocket/submarine"
	"github.com/gorilla/websocket"
	"github.com/olahol/melody"
)

// Session represents a network session that has user infos.
type Session struct {
	base *melody.Session
	id   uint64
	api  *api.WebSocketAPI
	room *Room
}

func newSession(melodySession *melody.Session, id uint64) *Session {
	serializer := typhenapi.NewJSONSerializer()
	session := &Session{melodySession, id, nil, nil}
	session.api = api.New(session, serializer, session.onError)
	session.api.Battle.OnPingReceive = session.onPingReceive
	return session
}

// Send sends raw message data to client.
func (session *Session) Send(msg []byte) {
	session.base.WriteBinary(msg)
}

func (session *Session) close() {
	session.base.Close()
}

func (session *Session) onMessage(data []byte) {
	session.api.DispatchMessageEvent(data)
}

func (session *Session) onError(data interface{}, err error) {
	if closeError, ok := err.(*websocket.CloseError); ok {
		if closeError.Code != 1000 {
			Log.Error(err)
		}
	}
}

func (session *Session) onPingReceive(message *battle.PingObject) {
	message.Message += " " + message.Message
	session.api.Battle.SendPing(message)
}
