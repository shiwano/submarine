package main_test

import (
	"app"
	"app/typhenapi/core"
	websocketapi "app/typhenapi/websocket/submarine"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http/httptest"
	"strings"
)

type clientSession struct {
	conn *websocket.Conn
	api  *websocketapi.WebSocketAPI
}

func (session *clientSession) Send(msg []byte) {
	session.conn.WriteMessage(websocket.BinaryMessage, msg)
}

func (session *clientSession) close() {
	session.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	session.conn.Close()
}

func (session *clientSession) readMessage() error {
	_, data, err := session.conn.ReadMessage()
	if err != nil {
		return err
	}
	session.api.DispatchMessageEvent(data)
	return nil
}

func newClientSession(url string) (*clientSession, error) {
	dialer := new(websocket.Dialer)
	conn, _, connErr := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	if connErr != nil {
		return nil, connErr
	}
	serializer := typhenapi.NewJSONSerializer()
	session := new(clientSession)
	session.conn = conn
	session.api = websocketapi.New(session, serializer, nil)
	return session, nil
}

func newTestServer() (*httptest.Server, *main.Server) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	rawServer := main.NewServer()
	server := httptest.NewServer(rawServer)
	return server, rawServer
}
