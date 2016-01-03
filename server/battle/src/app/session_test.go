package main_test

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSession(t *testing.T) {
	Convey("Session", t, func() {
		server, rawServer := newTestServer()
		session := newClientSession()
		session.connect(server.URL + "/rooms/1?room_key=key_1")

		Convey("should respond to a ping message", func() {
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
