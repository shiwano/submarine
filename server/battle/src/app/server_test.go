package main_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server := newTestServer()
		s := newClientSession()

		Convey("with valid params", func() {
			Convey("should connect", func() {
				err := s.connect(server.URL + "/rooms/1?room_key=key_1")
				So(err, ShouldBeNil)
			})
		})

		Convey("with a invalid room key", func() {
			Convey("should not connect", func() {
				err := s.connect(server.URL + "/rooms/1?room_key=invalid_room_key")
				So(err, ShouldNotBeNil)
			})
		})

		Convey("with a invalid room id", func() {
			Convey("should not connect", func() {
				err := s.connect(server.URL + "/rooms/400?room_key=key_1")
				So(err, ShouldNotBeNil)
			})
		})

		Reset(func() {
			server.Close()
		})
	})
}
