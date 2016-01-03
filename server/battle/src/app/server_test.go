package main_test

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server, rawServer := newTestServer()
		session := newClientSession()

		Convey("with valid a room key", func() {
			connectionError := session.connect(server.URL + "/rooms/1?room_key=key_1")

			Convey("should connect", func() {
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
			})
		})

		Convey("with invalid a room key", func() {
			connectionError := session.connect(server.URL + "/rooms/1?room_key=invalid_room_key")

			Convey("should not connect", func() {
				So(connectionError, ShouldNotBeNil)
			})

			Reset(func() {
				rawServer.Close()
			})
		})

		Convey("with invalid a room id", func() {
			connectionError := session.connect(server.URL + "/rooms/400?room_key=key_1")

			Convey("should not connect", func() {
				So(connectionError, ShouldNotBeNil)
			})

			Reset(func() {
				rawServer.Close()
			})
		})

		Reset(func() {
			server.Close()
		})
	})
}
