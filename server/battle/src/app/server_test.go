package main_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server := newTestServer()
		s := newClientSession()

		Convey("with valid params", func() {
			connError := s.connect(server.URL + "/rooms/1?room_key=key_1")

			Convey("should connect", func() {
				So(connError, ShouldBeNil)
			})
		})

		Convey("with a invalid room key", func() {
			connError := s.connect(server.URL + "/rooms/1?room_key=invalid_room_key")

			Convey("should not connect", func() {
				So(connError, ShouldNotBeNil)
			})
		})

		Convey("with a invalid room id", func() {
			connError := s.connect(server.URL + "/rooms/400?room_key=key_1")

			Convey("should not connect", func() {
				So(connError, ShouldNotBeNil)
			})
		})

		Reset(func() {
			server.Close()
		})
	})
}
