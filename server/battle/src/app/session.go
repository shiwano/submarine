package main

import (
	"app/connection"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	api "app/typhenapi/websocket/submarine"
	"github.com/gorilla/websocket"
	"net/http"
)

// Session represents a network session that has user infos.
type Session struct {
	conn *connection.Connection
	id   uint64
	api  *api.WebSocketAPI
	room *Room
}

func newSession() *Session {
	serializer := typhenapi.NewJSONSerializer()
	session := &Session{id: 1}

	session.api = api.New(session, serializer, session.onAPIError)
	session.api.Battle.OnPingReceive = session.onPingReceive

	session.conn = connection.NewConnection(connection.NewSettings())
	session.conn.OnMessageReceive = session.onConnectionMessageReceive
	session.conn.OnDisconnect = session.onConnectionDisconnect
	session.conn.OnError = session.onConnectionError
	return session
}

// Connect connects to the client.
func (session *Session) Connect(responseWriter http.ResponseWriter, request *http.Request) error {
	return session.conn.Connect(responseWriter, request)
}

// Send sends a binary message to the client.
func (session *Session) Send(data []byte) {
	session.conn.WriteBinaryMessage <- data
}

func (session *Session) close() {
	session.conn.WriteCloseMessage <- struct{}{}
}

func (session *Session) onConnectionDisconnect() {
	session.room.leave(session)
	session.room = nil
}

func (session *Session) onConnectionMessageReceive(data []byte) {
	session.api.DispatchMessageEvent(data)
}

func (session *Session) onConnectionError(err error) {
	if closeError, ok := err.(*websocket.CloseError); ok {
		if closeError.Code != 1000 {
			Log.Error(session.id, err)
		}
	}
}

func (session *Session) onAPIError(data interface{}, err error) {
	Log.Error(session.id, err)
}

func (session *Session) onPingReceive(message *battle.PingObject) {
	message.Message += " " + message.Message
	session.api.Battle.SendPing(message)
}
