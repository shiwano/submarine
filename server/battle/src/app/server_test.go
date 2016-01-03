package main_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server, rawServer := newTestServer()
		session := newClientSession()

		Convey("with valid params", func() {
			connectionError := session.connect(server.URL + "/rooms/1?room_key=key_1")

			Convey("should connect", func() {
				So(connectionError, ShouldBeNil)
			})

			Reset(func() {
				rawServer.Close()
				<-session.disconnected
			})
		})

		Convey("with a invalid room key", func() {
			connectionError := session.connect(server.URL + "/rooms/1?room_key=invalid_room_key")

			Convey("should not connect", func() {
				So(connectionError, ShouldNotBeNil)
			})

			Reset(func() {
				rawServer.Close()
			})
		})

		Convey("with a invalid room id", func() {
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
