package server

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/lib/typhenapi/core"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	webAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	webAPI_battle "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine/battle"
)

func TestWebAPI(t *testing.T) {
	Convey("WebAPI", t, func() {
		api := newWebAPIMock("http://localhost:3000")

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

func (m *webAPITransporter) Routes(path string, data []byte) (typhenapi.Type, int) {
	switch path {
	case "/ping":
		params := new(webAPI.PingRequestBody)
		m.serializer.Deserialize(data, params)
		return m.Ping(params)
	case "/battle/find_room":
		params := new(webAPI_battle.FindRoomRequestBody)
		m.serializer.Deserialize(data, params)
		return m.FindRoom(params)
	case "/battle/find_room_member":
		params := new(webAPI_battle.FindRoomMemberRequestBody)
		m.serializer.Deserialize(data, params)
		return m.FindRoomMember(params)
	case "/battle/close_room":
		params := new(webAPI_battle.CloseRoomRequestBody)
		m.serializer.Deserialize(data, params)
		return m.CloseRoom(params)
	default:
		return new(typhenapi.Void), http.StatusNotFound
	}
}

func (m *webAPITransporter) Ping(params *webAPI.PingRequestBody) (typhenapi.Type, int) {
	typhenType := &api.PingObject{Message: params.Message + " PONG"}
	return typhenType, http.StatusOK
}

func (m *webAPITransporter) FindRoomMember(params *webAPI_battle.FindRoomMemberRequestBody) (typhenapi.Type, int) {
	typhenType := new(battleAPI.FindRoomMemberObject)

	switch params.RoomKey {
	case "key_1":
		typhenType.RoomMember = &battleAPI.RoomMember{Id: 1, RoomId: 1, Name: "I168"}
	case "key_2":
		typhenType.RoomMember = &battleAPI.RoomMember{Id: 2, RoomId: 1, Name: "I8"}
	case "key_3":
		typhenType.RoomMember = &battleAPI.RoomMember{Id: 3, RoomId: 1, Name: "I19"}
	case "key_4":
		typhenType.RoomMember = &battleAPI.RoomMember{Id: 4, RoomId: 1, Name: "I58"}
	}
	return typhenType, http.StatusOK
}

func (m *webAPITransporter) FindRoom(params *webAPI_battle.FindRoomRequestBody) (typhenapi.Type, int) {
	typhenType := new(battleAPI.FindRoomObject)

	if params.RoomId == 1 {
		typhenType.Room = &battleAPI.Room{
			Id: params.RoomId,
		}
	}
	return typhenType, http.StatusOK
}

func (m *webAPITransporter) CloseRoom(params *webAPI_battle.CloseRoomRequestBody) (typhenapi.Type, int) {
	if params.RoomId <= 0 {
		return new(typhenapi.Void), http.StatusInternalServerError
	}
	return new(typhenapi.Void), http.StatusOK
}
