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
	id     uint64
	roomID uint64
	conn   *connection.Connection
	api    *api.WebSocketAPI
	room   *Room
}

func newSession(id uint64, roomID uint64) *Session {
	serializer := typhenapi.NewJSONSerializer()
	session := &Session{id: id, roomID: roomID}

	session.api = api.New(session, serializer, session.onError)
	session.api.Battle.PingHandler = session.onPingReceive

	session.conn = connection.New()
	session.conn.OnBinaryMessageReceive = session.onConnectionBinaryMessageReceive
	session.conn.OnDisconnect = session.onConnectionDisconnect
	session.conn.OnError = session.onError
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
	session.conn.Close()
}

func (session *Session) onConnectionDisconnect() {
	if session.room != nil {
		session.room.Leave <- session
	}
}

func (session *Session) onConnectionBinaryMessageReceive(data []byte) {
	session.api.DispatchMessageEvent(data)
}

func (session *Session) onError(err error) {
	if closeError, ok := err.(*websocket.CloseError); ok && closeError.Code == 1000 {
		return
	}
	Log.Error(session.id, err)
}

func (session *Session) onPingReceive(message *battle.PingObject) {
	message.Message += " " + message.Message
	session.api.Battle.SendPing(message)
}
