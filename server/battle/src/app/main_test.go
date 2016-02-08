package main_test

import (
	"app"
	"app/conn"
	"app/logger"
	"app/typhenapi/core"
	webapi "app/typhenapi/web/submarine"
	websocketapi "app/typhenapi/websocket/submarine"
	"bytes"
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var p = pp.Println

type clientSession struct {
	conn         *conn.Conn
	api          *websocketapi.WebSocketAPI
	disconnected chan struct{}
}

func newClientSession() *clientSession {
	serializer := typhenapi.NewJSONSerializer()
	session := &clientSession{
		conn: conn.New(),
	}
	session.api = websocketapi.New(session, serializer, nil)
	session.conn.DisconnectHandler = func() {
		session.disconnected <- struct{}{}
	}
	session.conn.BinaryMessageHandler = func(data []byte) {
		if err := session.api.DispatchMessageEvent(data); err != nil {
			logger.Log.Error(err)
		}
	}
	return session
}

func (s *clientSession) Send(data []byte) error {
	return s.conn.WriteBinaryMessage(data)
}

func (s *clientSession) connect(url string) error {
	s.disconnected = make(chan struct{}, 1)
	_, err := s.conn.Connect(strings.Replace(url, "http", "ws", 1), nil)
	return err
}

func (s *clientSession) close() {
	s.conn.Close()
}

func (s *clientSession) waitForDisconnected() {
	if s.disconnected != nil {
		<-s.disconnected
	}
}

type webAPITransporter struct {
	serializer typhenapi.Serializer
}

func (m *webAPITransporter) RoundTrip(request *http.Request) (*http.Response, error) {
	response := &http.Response{Header: make(http.Header), Request: request}
	response.Header.Set("Content-Type", "application/json")
	data, _ := ioutil.ReadAll(request.Body)
	typhenType, statusCode := m.Routes(request.URL.Path, data)
	err := typhenType.Coerce()
	if err != nil {
		return response, err
	}

	response.StatusCode = statusCode
	body, _ := typhenType.Bytes(m.serializer)
	response.Body = ioutil.NopCloser(bytes.NewReader(body))
	if response.StatusCode >= 400 {
		return nil, errors.New("Server Error")
	}
	return response, nil
}

func newWebAPIMock(url string) *webapi.WebAPI {
	main.WebAPIRoundTripper = &webAPITransporter{typhenapi.NewJSONSerializer()}
	return main.NewWebAPI(url)
}

func newTestServer() *httptest.Server {
	logger.Log.Level = logrus.PanicLevel
	main.WebAPIRoundTripper = &webAPITransporter{typhenapi.NewJSONSerializer()}
	gin.SetMode(gin.TestMode)
	server := httptest.NewServer(main.NewServer())
	return server
}
