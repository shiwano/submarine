package server

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

func TestRoom(t *testing.T) {
	Convey("Room", t, func() {
		server := newTestServer()
		s := newClientSession()
		s2 := newClientSession()

		Convey("when a user join to the room", func() {
			Convey("should send a room message", func() {
				done := make(chan *api.Room)
				s.api.Battle.RoomHandler = func(m *api.Room) { done <- m }
				s.connect(server.URL + "/rooms/1?room_key=key_1")
				m := <-done
				So(m.Id, ShouldEqual, 1)
				So(m.Members, ShouldHaveLength, 1)
			})

			Convey("should send a now message", func() {
				currentmillis.StubNow = func() int64 { return 123456 }
				done := make(chan *battleAPI.NowObject)
				s.api.Battle.NowHandler = func(m *battleAPI.NowObject) { done <- m }
				s.connect(server.URL + "/rooms/1?room_key=key_1")
				m := <-done
				currentmillis.StubNow = nil
				So(m.Time, ShouldEqual, 123456)
			})
		})

		Convey("when a new user join to the room", func() {
			Convey("should send a room message to other users", func() {
				done := make(chan *api.Room)
				s.api.Battle.RoomHandler = func(m *api.Room) { done <- m }
				s.connect(server.URL + "/rooms/1?room_key=key_1")
				s2.connect(server.URL + "/rooms/1?room_key=key_2")
				<-done
				m := <-done
				So(m.Id, ShouldEqual, 1)
				So(m.Members, ShouldHaveLength, 2)
			})
		})

		Reset(func() {
			server.Close()
		})
	})
}
