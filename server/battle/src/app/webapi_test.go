package main_test

import (
	"app/typhenapi/core"
	"app/typhenapi/type/submarine"
	webapi "app/typhenapi/web/submarine"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestWebAPI(t *testing.T) {
	Convey("WebAPI", t, func() {
		api := newWebAPIMock("http://localhost:3000")

		Convey("should send a ping request", func() {
			res, _ := api.Ping("PING")
			So(res.Message, ShouldEqual, "PING PONG")
		})
	})
}

func (m *webAPITransporter) Routes(path string, data []byte) (typhenapi.Type, int) {
	switch path {
	case "/ping":
		params := new(webapi.PingRequestBody)
		m.serializer.Deserialize(data, params)
		return m.Ping(params)
	default:
		return &typhenapi.Void{}, http.StatusOK
	}
}

func (m *webAPITransporter) Ping(params *webapi.PingRequestBody) (typhenapi.Type, int) {
	typhenType := &submarine.PingObject{"PING PONG"}
	return typhenType, http.StatusOK
}
