package main_test

import (
	"app"
	"app/connection"
	"app/typhenapi/core"
	webapi "app/typhenapi/web/submarine"
	websocketapi "app/typhenapi/websocket/submarine"
	"bytes"
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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
	s.conn.WriteBinaryMessage(data)
}

func (s *clientSession) close() {
	s.conn.Close()
}

type webAPITransporter struct {
	serializer typhenapi.Serializer
}

func (m *webAPITransporter) RoundTrip(request *http.Request) (*http.Response, error) {
	response := &http.Response{Header: make(http.Header), Request: request}
	response.Header.Set("Content-Type", "application/json")
	data, _ := ioutil.ReadAll(request.Body)
	typhenType, statusCode := m.Routes(request.URL.Path, data)

	response.StatusCode = statusCode
	body, _ := typhenType.Bytes(m.serializer)
	response.Body = ioutil.NopCloser(bytes.NewReader(body))
	if response.StatusCode >= 400 {
		return response, errors.New("Server Error")
	}
	return response, nil
}

func newWebAPIMock(url string) *webapi.WebAPI {
	api := main.NewWebAPI(url)
	api.Client.Transport = &webAPITransporter{typhenapi.NewJSONSerializer()}
	return api
}

func newTestServer() (*httptest.Server, *main.Server) {
	main.Log.Level = logrus.WarnLevel
	gin.SetMode(gin.TestMode)
	rawServer := main.NewServer()
	server := httptest.NewServer(rawServer)
	return server, rawServer
}
