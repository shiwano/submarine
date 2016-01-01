package main_test

import (
	"app"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine"
	webapi "app/typhenapi/web/submarine"
	"bytes"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
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

func (mock *webAPITransporter) Routes(path string, data []byte) (typhenapi.Type, int) {
	switch path {
	case "/ping":
		params := new(webapi.PingRequestBody)
		mock.serializer.Deserialize(data, params)
		return mock.Ping(params)
	default:
		return &typhenapi.Void{}, http.StatusOK
	}
}

func (mock *webAPITransporter) Ping(params *webapi.PingRequestBody) (typhenapi.Type, int) {
	typhenType := &submarine.PingObject{"PING PONG"}
	return typhenType, http.StatusOK
}
