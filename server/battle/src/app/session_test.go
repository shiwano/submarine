package main_test

import (
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSession(t *testing.T) {
	Convey("Session", t, func() {
		server, rawServer := newTestServer()
		session := newClientSession()
		session.connect(server.URL + "/rooms/1?room_key=key_1")

		Convey("should respond/receive to a ping message", func() {
			done := make(chan *battle.PingObject)
			session.api.Battle.PingHandler = func(m *battle.PingObject) { done <- m }
			session.api.Battle.SendPing(&battle.PingObject{"Hey"})
			m := <-done
			So(m.Message, ShouldEqual, "Hey Hey")
		})

		Convey("should respond to a room message", func() {
			done := make(chan *submarine.Room)
			session.api.Battle.RoomHandler = func(m *submarine.Room) { done <- m }
			m := <-done
			So(m.Id, ShouldEqual, 1)
			So(m.Members, ShouldHaveLength, 1)
		})

		Reset(func() {
			rawServer.Close()
			<-session.disconnected
			server.Close()
		})
	})
}
