package main_test

import (
	"app"
	"app/connection"
	"app/typhenapi/core"
	websocketapi "app/typhenapi/websocket/submarine"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
)

type clientSession struct {
	conn         *connection.Connection
	api          *websocketapi.WebSocketAPI
	disconnected chan struct{}
}

func newClientSession(url string) (*clientSession, error) {
	conn := connection.New()
	_, err := conn.Connect(strings.Replace(url, "http", "ws", 1), nil)
	if err != nil {
		return nil, err
	}

	serializer := typhenapi.NewJSONSerializer()
	session := new(clientSession)
	session.disconnected = make(chan struct{})
	session.conn = conn
	session.conn.DisconnectHandler = func() {
		session.disconnected <- struct{}{}
	}
	session.conn.BinaryMessageHandler = func(data []byte) {
		if err := session.api.DispatchMessageEvent(data); err != nil {
			main.Log.Error(err)
		}
	}
	session.api = websocketapi.New(session, serializer, nil)
	return session, nil
}

func (s *clientSession) Send(data []byte) {
	s.conn.WriteBinaryMessage <- data
}

func (s *clientSession) close() {
	s.conn.Close()
}

func newTestServer() (*httptest.Server, *main.Server) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	rawServer := main.NewServer()
	server := httptest.NewServer(rawServer)
	return server, rawServer
}
