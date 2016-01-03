package main_test

import (
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSession(t *testing.T) {
	Convey("Session", t, func() {
		server := newTestServer()
		s := newClientSession()
		s2 := newClientSession()

		Convey("should respond/receive to a ping message", func() {
			done := make(chan *battle.PingObject)
			s.connect(server.URL + "/rooms/1?room_key=key_1")
			s.api.Battle.PingHandler = func(m *battle.PingObject) { done <- m }
			s.api.Battle.SendPing(&battle.PingObject{"Hey"})
			m := <-done
			So(m.Message, ShouldEqual, "Hey Hey")
		})

		Convey("should join to the room", func() {
			done := make(chan *submarine.Room)
			s2.api.Battle.RoomHandler = func(m *submarine.Room) { done <- m }
			s.connect(server.URL + "/rooms/1?room_key=key_1")
			s2.connect(server.URL + "/rooms/1?room_key=key_2")
			m := <-done
			So(m.Id, ShouldEqual, 1)
			So(m.Members, ShouldHaveLength, 2)
		})

		Reset(func() {
			server.Close()
		})
	})
}
