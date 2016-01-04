package main_test

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSession(t *testing.T) {
	Convey("Session", t, func() {
		server := newTestServer()
		s := newClientSession()

		Convey("should respond to a ping message", func() {
			done := make(chan *battle.PingObject)
			s.connect(server.URL + "/rooms/1?room_key=key_1")
			s.api.Battle.PingHandler = func(m *battle.PingObject) { done <- m }
			s.api.Battle.SendPing(&battle.PingObject{"Hey"})
			m := <-done
			So(m.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			server.Close()
		})
	})
}
