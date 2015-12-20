package main_test

import (
	"app"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	battleapi "app/typhenapi/websocket/submarine/battle"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer() (*httptest.Server, *io.PipeWriter) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	engine, logWriter := main.NewEngine()
	server := httptest.NewServer(engine)
	return server, logWriter
}

func newDialer(url string) (*websocket.Conn, error) {
	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	return conn, err
}

func TestNewEngine(t *testing.T) {
	server, logWriter := newTestServer()
	defer logWriter.Close()
	defer server.Close()

	conn, connErr := newDialer(server.URL + "/room/1")
	defer conn.Close()
	assert.NoError(t, connErr)

	serializer := typhenapi.NewJSONSerializer()
	ping := &battle.Ping{"Foobar"}
	message, messageErr := typhenapi.NewMessage(serializer, battleapi.MessageType_Ping, ping)
	assert.NoError(t, messageErr)
	conn.WriteMessage(websocket.BinaryMessage, message.Bytes())

	messageType, received, readMessageErr := conn.ReadMessage()
	assert.NoError(t, readMessageErr)
	assert.Equal(t, websocket.BinaryMessage, messageType)

	assert.Equal(t, message.Bytes(), received)
}
