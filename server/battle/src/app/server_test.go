package main_test

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server, rawServer := newTestServer()
		session, connectionError := newClientSession(server.URL + "/rooms/1?room_key=secret")

		Convey("should be connectable by websocket protocol", func() {
			So(connectionError, ShouldBeNil)
		})

		Convey("should respond to a ping websocket message", func() {
			done := make(chan *battle.PingObject)
			session.api.Battle.PingHandler = func(message *battle.PingObject) { done <- message }
			session.api.Battle.SendPing(&battle.PingObject{"Hey"})
			message := <-done
			So(message.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			rawServer.Close()
			<-session.disconnected
			server.Close()
		})
	})
}
