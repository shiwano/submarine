package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/shiwano/websocket-conn"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi/core"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	rtmAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/websocket/submarine"
	"github.com/shiwano/submarine/server/battle/server/logger"
)

type session struct {
	id                int64
	roomID            int64
	info              *battleAPI.RoomMember
	conn              *conn.Conn
	api               *rtmAPI.WebSocketAPI
	room              *room
	disconnectHandler func(*session)
}

func newSession(info *battleAPI.RoomMember, roomID int64) *session {
	serializer := typhenapi.NewJSONSerializer()
	s := &session{
		id:     info.Id,
		info:   info,
		roomID: roomID,
		conn:   conn.New(),
	}
	s.conn.BinaryMessageHandler = s.onBinaryMessageReceive
	s.conn.DisconnectHandler = s.onDisconnect
	s.conn.ErrorHandler = s.onError

	s.api = rtmAPI.New(s, serializer, s.onError)
	s.api.Battle.PingHandler = s.onPingReceive
	s.api.Battle.StartRequestHandler = s.onStartRequestReceive
	s.api.Battle.AddBotRequestHandler = s.onAddBotRequestReceive
	s.api.Battle.RemoveBotRequestHandler = s.onRemoveBotRequestReceive
	s.api.Battle.AccelerationRequestHandler = func(m *battleAPI.AccelerationRequestObject) { s.onBattleMessageReceive(m) }
	s.api.Battle.BrakeRequestHandler = func(m *battleAPI.BrakeRequestObject) { s.onBattleMessageReceive(m) }
	s.api.Battle.TurnRequestHandler = func(m *battleAPI.TurnRequestObject) { s.onBattleMessageReceive(m) }
	s.api.Battle.TorpedoRequestHandler = func(m *battleAPI.TorpedoRequestObject) { s.onBattleMessageReceive(m) }
	s.api.Battle.PingerRequestHandler = func(m *battleAPI.PingerRequestObject) { s.onBattleMessageReceive(m) }
	s.api.Battle.WatcherRequestHandler = func(m *battleAPI.WatcherRequestObject) { s.onBattleMessageReceive(m) }
	return s
}

func (s *session) String() string {
	return fmt.Sprintf("Session(%v)", s.id)
}

func (s *session) Connect(responseWriter http.ResponseWriter, request *http.Request) error {
	return s.conn.UpgradeFromHTTP(responseWriter, request)
}

func (s *session) Send(data []byte) error {
	return s.conn.WriteBinaryMessage(data)
}

func (s *session) close() {
	s.conn.Close()
}

func (s *session) toUserAPIType() *api.User {
	return &api.User{Name: s.info.Name}
}

func (s *session) onDisconnect() {
	if s.disconnectHandler != nil {
		s.disconnectHandler(s)
	}
}

func (s *session) onBinaryMessageReceive(data []byte) {
	s.api.DispatchMessageEvent(data)
}

func (s *session) onError(err error) {
	if closeError, ok := err.(*websocket.CloseError); ok && (closeError.Code == 1000 || closeError.Code == 1005) {
		return
	}
	logger.Log.Errorf("%v received the error: %v", s, err)
}

func (s *session) onPingReceive(m *battleAPI.PingObject) {
	m.Message += " " + m.Message
	s.api.Battle.SendPing(m)
}

func (s *session) onStartRequestReceive(m *battleAPI.StartRequestObject) {
	s.room.startBattleCh <- s
}

func (s *session) onAddBotRequestReceive(m *battleAPI.AddBotRequestObject) {
	s.room.addBotCh <- struct{}{}
}

func (s *session) onRemoveBotRequestReceive(m *battleAPI.RemoveBotRequestObject) {
	s.room.removeBotCh <- m.BotId
}

func (s *session) onBattleMessageReceive(m typhenapi.Type) {
	if s.room != nil && !s.room.isClosed.IsSet() {
		s.room.sendBattleInput(s.id, m)
	}
}

func (s *session) synchronizeTime() {
	s.api.Battle.SendNow(&battleAPI.NowObject{Time: currentmillis.Now()})
}
