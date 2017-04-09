package server

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
	"github.com/shiwano/submarine/server/battle/src/config"
)

func TestWebAPI(t *testing.T) {
	Convey("WebAPI", t, func() {
		api := newWebAPIMock(config.Config.ApiServerBaseUri)

		Convey("should send a ping request", func() {
			res, err := api.Ping("PING")
			So(err, ShouldBeNil)
			So(res.Message, ShouldEqual, "PING PONG")
		})

		Convey("should send a battle/close_room request", func() {
			res, err := api.Battle.CloseRoom(1)
			So(err, ShouldBeNil)
			So(res, ShouldHaveSameTypeAs, new(typhenapi.Void))
		})

		Convey("should send a battle/find_room request", func() {
			res, err := api.Battle.FindRoom(1)
			So(err, ShouldBeNil)
			So(res.Room.Id, ShouldEqual, 1)
		})
	})
}
