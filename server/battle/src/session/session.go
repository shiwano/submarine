package session

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	rtmAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/websocket/submarine"
	"github.com/shiwano/submarine/server/battle/src/logger"
	"github.com/shiwano/websocket-conn"
)

var connectionSettings = conn.DefaultSettings()

// Session represents a session of the battle server.
type Session struct {
	roomMember            *battleAPI.RoomMember
	api                   *rtmAPI.WebSocketAPI
	conn                  *conn.Conn
	roomMessageReceived   chan typhenapi.Type
	battleMessageReceived chan typhenapi.Type
	closed                chan struct{}
}

// New creates a Session.
func New(ctx context.Context, roomMember *battleAPI.RoomMember, responseWriter http.ResponseWriter, request *http.Request) (*Session, error) {
	c, err := conn.UpgradeFromHTTP(ctx, connectionSettings, responseWriter, request)
	if err != nil {
		return nil, err
	}

	s := &Session{
		roomMember:            roomMember,
		conn:                  c,
		roomMessageReceived:   make(chan typhenapi.Type),
		battleMessageReceived: make(chan typhenapi.Type),
		closed:                make(chan struct{}),
	}
	s.api = rtmAPI.New(s, new(typhenapi.MessagePackSerializer), s.onError)
	s.api.Battle.PingHandler = s.onPingReceive
	s.api.Battle.StartRequestHandler = func(m *battleAPI.StartRequest) { s.onRoomMessageReceive(m) }
	s.api.Battle.AddBotRequestHandler = func(m *battleAPI.AddBotRequest) { s.onRoomMessageReceive(m) }
	s.api.Battle.RemoveBotRequestHandler = func(m *battleAPI.RemoveBotRequest) { s.onRoomMessageReceive(m) }
	s.api.Battle.AccelerationRequestHandler = func(m *battleAPI.AccelerationRequest) { s.onBattleMessageReceive(m) }
	s.api.Battle.BrakeRequestHandler = func(m *battleAPI.BrakeRequest) { s.onBattleMessageReceive(m) }
	s.api.Battle.TurnRequestHandler = func(m *battleAPI.TurnRequest) { s.onBattleMessageReceive(m) }
	s.api.Battle.TorpedoRequestHandler = func(m *battleAPI.TorpedoRequest) { s.onBattleMessageReceive(m) }
	s.api.Battle.PingerRequestHandler = func(m *battleAPI.PingerRequest) { s.onBattleMessageReceive(m) }
	s.api.Battle.WatcherRequestHandler = func(m *battleAPI.WatcherRequest) { s.onBattleMessageReceive(m) }

	go s.run()
	return s, nil
}

func (s *Session) String() string {
	return fmt.Sprintf("Session(%v)", s.roomMember.Id)
}

// Send implements typhenapi.Type.
func (s *Session) Send(data []byte) error {
	return s.conn.SendBinaryMessage(data)
}

// ID returns the session ID.
func (s *Session) ID() int64 {
	return s.roomMember.Id
}

// RoomMember returns the session's room member data.
func (s *Session) RoomMember() *battleAPI.RoomMember {
	return s.roomMember
}

// ToUserAPIType converts the session to API's User type.
func (s *Session) ToUserAPIType() *api.User {
	return &api.User{Name: s.roomMember.Name}
}

// SynchronizeTime sends Now message to session for time sync.
func (s *Session) SynchronizeTime() {
	s.api.Battle.SendNow(&battleAPI.Now{Time: currentmillis.Now()})
}

// SendRoom sends Room message.
func (s *Session) SendRoom(room *battleAPI.Room) {
	s.api.Battle.SendRoom(room)
}

// SendBattleMessage sends a message that is defined in Battle namespace.
func (s *Session) SendBattleMessage(m typhenapi.Type) {
	s.api.Battle.Send(m)
}

// RoomMessageReceived returns a channel that receives a value when a room message was received.
func (s *Session) RoomMessageReceived() <-chan typhenapi.Type {
	return s.roomMessageReceived
}

// BattleMessageReceived returns a channel that receives a value when a battle message was received.
func (s *Session) BattleMessageReceived() <-chan typhenapi.Type {
	return s.battleMessageReceived
}

// Closed returns a channel that receives a value when the session closed.
func (s *Session) Closed() <-chan struct{} {
	return s.closed
}

func (s *Session) run() {
	for m := range s.conn.Stream() {
		if m.MessageType == conn.BinaryMessageType {
			s.onBinaryMessageReceive(m.Data)
		}
	}
	s.onError(s.conn.Err())
	close(s.battleMessageReceived)
	close(s.roomMessageReceived)
	close(s.closed)
}

func (s *Session) onBinaryMessageReceive(data []byte) {
	s.api.DispatchMessageEvent(data)
}

func (s *Session) onError(err error) {
	if err == nil || err == context.Canceled {
		return
	}
	if closeError, ok := err.(*websocket.CloseError); ok {
		if closeError.Code == 1000 || closeError.Code == 1005 {
			return
		}
		logger.Log.Warnf("%v received the websocket close error: %v", s, closeError)
	} else {
		logger.Log.Errorf("%v received the error: %v", s, err)
	}
}

func (s *Session) onRoomMessageReceive(m typhenapi.Type) {
	s.roomMessageReceived <- m
}

func (s *Session) onBattleMessageReceive(m typhenapi.Type) {
	s.battleMessageReceived <- m
}

func (s *Session) onPingReceive(m *battleAPI.Ping) {
	s.api.Battle.SendPing(&battleAPI.Ping{
		Message: m.Message + " " + m.Message,
	})
}
