package main

import (
	"app/logger"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	api "app/typhenapi/websocket/submarine"
	"github.com/gorilla/websocket"
	"lib/conn"
	"lib/currentmillis"
	"net/http"
)

// Session represents a network session that has user infos.
type Session struct {
	id                int64
	roomID            int64
	info              *battle.RoomMember
	conn              *conn.Conn
	api               *api.WebSocketAPI
	room              *Room
	disconnectHandler func(*Session)
}

func newSession(info *battle.RoomMember, roomID int64) *Session {
	serializer := typhenapi.NewJSONSerializer()
	session := &Session{
		id:     info.Id,
		info:   info,
		roomID: roomID,
		conn:   conn.New(),
	}
	session.conn.BinaryMessageHandler = session.onBinaryMessageReceive
	session.conn.DisconnectHandler = session.onDisconnect
	session.conn.ErrorHandler = session.onError

	session.api = api.New(session, serializer, session.onError)
	session.api.Battle.PingHandler = session.onPingReceive
	session.api.Battle.AccelerationRequestHandler = session.onAccelerationRequestReceive
	session.api.Battle.BrakeRequestHandler = session.onBrakeRequestReceive
	session.api.Battle.TurnRequestHandler = session.onTurnRequestReceive
	session.api.Battle.TorpedoRequestHandler = session.onTorpedoRequestReceive
	session.api.Battle.PingerRequestHandler = session.onPingerRequestReceive
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

func (s *Session) toUserAPIType() *submarine.User {
	return &submarine.User{Name: s.info.Name}
}

func (s *Session) onDisconnect() {
	if s.disconnectHandler != nil {
		s.disconnectHandler(s)
	}
}

func (s *Session) onBinaryMessageReceive(data []byte) {
	s.api.DispatchMessageEvent(data)
}

func (s *Session) onError(err error) {
	if closeError, ok := err.(*websocket.CloseError); ok && (closeError.Code == 1000 || closeError.Code == 1005) {
		return
	}
	logger.Log.Error(s.id, err)
}

func (s *Session) onPingReceive(message *battle.PingObject) {
	message.Message += " " + message.Message
	s.api.Battle.SendPing(message)
}

func (s *Session) onAccelerationRequestReceive(message *battle.AccelerationRequestObject) {
	s.onBattleMessageReceive(message)
}

func (s *Session) onBrakeRequestReceive(message *battle.BrakeRequestObject) {
	s.onBattleMessageReceive(message)
}

func (s *Session) onTurnRequestReceive(message *battle.TurnRequestObject) {
	s.onBattleMessageReceive(message)
}

func (s *Session) onPingerRequestReceive(message *battle.PingerRequestObject) {
	s.onBattleMessageReceive(message)
}

func (s *Session) onTorpedoRequestReceive(message *battle.TorpedoRequestObject) {
	s.onBattleMessageReceive(message)
}

func (s *Session) onBattleMessageReceive(message typhenapi.Type) {
	if s.room != nil && !s.room.isClosed {
		s.room.sendBattleInput(s.id, message)
	}
}

func (s *Session) synchronizeTime() {
	s.api.Battle.SendNow(&battle.NowObject{Time: currentmillis.Now()})
}
