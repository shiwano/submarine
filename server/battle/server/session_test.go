package server

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

func TestSession(t *testing.T) {
	Convey("Session", t, func() {
		server := newTestServer()
		s := newClientSession()

		Convey("should respond to a ping message", func() {
			done := make(chan *battleAPI.PingObject)
			s.connect(server.URL + "/rooms/1?room_key=key_1")
			s.api.Battle.PingHandler = func(m *battleAPI.PingObject) { done <- m }
			s.api.Battle.SendPing(&battleAPI.PingObject{"Hey"})
			m := <-done
			So(m.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			server.Close()
		})
	})
}
