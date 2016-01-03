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
	id                int64
	roomID            int64
	conn              *connection.Connection
	api               *api.WebSocketAPI
	room              *Room
	disconnectHandler func(*Session)
}

func newSession(id int64, roomID int64) *Session {
	serializer := typhenapi.NewJSONSerializer()
	session := &Session{id: id, roomID: roomID}

	session.api = api.New(session, serializer, session.onError)
	session.api.Battle.PingHandler = session.onPingReceive

	session.conn = connection.New()
	session.conn.BinaryMessageHandler = session.onConnectionBinaryMessageReceive
	session.conn.DisconnectHandler = session.onConnectionDisconnect
	session.conn.ErrorHandler = session.onError
	return session
}

// Connect connects to the client.
func (s *Session) Connect(responseWriter http.ResponseWriter, request *http.Request) error {
	return s.conn.UpgradeFromHTTP(responseWriter, request)
}

// Send sends a binary message to the client.
func (s *Session) Send(data []byte) error {
	return s.conn.WriteBinaryMessage(data)
}

func (s *Session) close() {
	s.conn.Close()
}

func (s *Session) onConnectionDisconnect() {
	if s.disconnectHandler != nil {
		s.disconnectHandler(s)
	}
}

func (s *Session) onConnectionBinaryMessageReceive(data []byte) {
	s.api.DispatchMessageEvent(data)
}

func (s *Session) onError(err error) {
	if closeError, ok := err.(*websocket.CloseError); ok && (closeError.Code == 1000 || closeError.Code == 1005) {
		return
	}
	Log.Error(s.id, err)
}

func (s *Session) onPingReceive(message *battle.PingObject) {
	message.Message += " " + message.Message
	s.api.Battle.SendPing(message)
}
