package main_test

import (
	"app"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	api "app/typhenapi/websocket/submarine"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

type clientSession struct {
	*websocket.Conn
	api *api.WebSocketAPI
}

func (session *clientSession) Send(msg []byte) {
	session.WriteMessage(websocket.BinaryMessage, msg)
}

func (session *clientSession) readMessage() error {
	_, data, err := session.ReadMessage()
	if err != nil {
		return err
	}
	session.api.DispatchMessageEvent(data)
	return nil
}

func newDialer(url string) (*websocket.Conn, error) {
	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	return conn, err
}

func newSession(url string) (*clientSession, error) {
	conn, connErr := newDialer(url)
	if connErr != nil {
		return nil, connErr
	}
	serializer := typhenapi.NewJSONSerializer()
	session := &clientSession{conn, nil}
	session.api = api.New(session, serializer, nil)
	return session, nil
}

func newTestServer() (*httptest.Server, *io.PipeWriter) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	engine, logWriter := main.NewEngine()
	server := httptest.NewServer(engine)
	return server, logWriter
}

func TestBattleServer(t *testing.T) {
	Convey("BattleServer", t, func() {
		server, logWriter := newTestServer()

		Convey("should be connectable by web socket protocol", func() {
			done := make(chan error)
			go func() {
				conn, err := newDialer(server.URL + "/room/1")
				defer conn.Close()
				done <- err
			}()
			err := <-done
			So(err, ShouldBeNil)
		})

		Convey("should respond to a ping message", func() {
			done := make(chan *battle.Ping)
			go func() {
				session, _ := newSession(server.URL + "/room/1")
				defer session.Close()
				session.api.Battle.OnPingReceive = func(message *battle.Ping) { done <- message }
				session.api.Battle.SendPing(&battle.Ping{"Hey"})
				session.readMessage()
			}()
			message := <-done
			So(message.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			server.Close()
			logWriter.Close()
		})
	})
}
