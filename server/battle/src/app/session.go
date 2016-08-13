package main

import (
	"app/logger"
	"app/typhenapi/core"
	api "app/typhenapi/type/submarine"
	battleAPI "app/typhenapi/type/submarine/battle"
	rtmAPI "app/typhenapi/websocket/submarine"
	"github.com/gorilla/websocket"
	"lib/conn"
	"lib/currentmillis"
	"net/http"
)

// Session represents a network session that has user infos.
type Session struct {
	id                int64
	roomID            int64
	info              *battleAPI.RoomMember
	conn              *conn.Conn
	api               *rtmAPI.WebSocketAPI
	room              *Room
	disconnectHandler func(*Session)
}

func newSession(info *battleAPI.RoomMember, roomID int64) *Session {
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

	session.api = rtmAPI.New(session, serializer, session.onError)
	session.api.Battle.PingHandler = session.onPingReceive
	session.api.Battle.StartRequestHandler = session.onStartRequestReceive
	session.api.Battle.AddBotRequestHandler = session.onAddBotRequestReceive
	session.api.Battle.RemoveBotRequestHandler = session.onRemoveBotRequestReceive
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

func (s *Session) toUserAPIType() *api.User {
	return &api.User{Name: s.info.Name}
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

func (s *Session) onPingReceive(m *battleAPI.PingObject) {
	m.Message += " " + m.Message
	s.api.Battle.SendPing(m)
}

func (s *Session) onStartRequestReceive(m *battleAPI.StartRequestObject) {
	s.room.startBattleCh <- s
}

func (s *Session) onAddBotRequestReceive(m *battleAPI.AddBotRequestObject) {
	s.room.addBotCh <- struct{}{}
}

func (s *Session) onRemoveBotRequestReceive(m *battleAPI.RemoveBotRequestObject) {
	s.room.removeBotCh <- m.BotId
}

func (s *Session) onAccelerationRequestReceive(m *battleAPI.AccelerationRequestObject) {
	s.onBattleMessageReceive(m)
}

func (s *Session) onBrakeRequestReceive(m *battleAPI.BrakeRequestObject) {
	s.onBattleMessageReceive(m)
}

func (s *Session) onTurnRequestReceive(m *battleAPI.TurnRequestObject) {
	s.onBattleMessageReceive(m)
}

func (s *Session) onPingerRequestReceive(m *battleAPI.PingerRequestObject) {
	s.onBattleMessageReceive(m)
}

func (s *Session) onTorpedoRequestReceive(m *battleAPI.TorpedoRequestObject) {
	s.onBattleMessageReceive(m)
}

func (s *Session) onBattleMessageReceive(m typhenapi.Type) {
	if s.room != nil && !s.room.isClosed.IsSet() {
		s.room.sendBattleInput(s.id, m)
	}
}

func (s *Session) synchronizeTime() {
	s.api.Battle.SendNow(&battleAPI.NowObject{Time: currentmillis.Now()})
}
