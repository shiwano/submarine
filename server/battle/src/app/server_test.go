package main_test

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {
	Convey("Server", t, func() {
		server, rawServer := newTestServer()

		Convey("should be connectable by websocket protocol", func() {
			done := make(chan error)
			go func() {
				session, err := newClientSession(server.URL + "/rooms/1?room_key=secret")
				defer session.close()
				done <- err
			}()
			err := <-done
			So(err, ShouldBeNil)
		})

		Convey("should respond to a ping websocket message", func() {
			done := make(chan *battle.PingObject)
			go func() {
				session, _ := newClientSession(server.URL + "/rooms/1?room_key=secret")
				defer session.close()
				session.api.Battle.PingHandler = func(message *battle.PingObject) { done <- message }
				session.api.Battle.SendPing(&battle.PingObject{"Hey"})
				session.readMessage()
			}()
			message := <-done
			So(message.Message, ShouldEqual, "Hey Hey")
		})

		Reset(func() {
			server.Close()
			rawServer.Close()
		})
	})
}
