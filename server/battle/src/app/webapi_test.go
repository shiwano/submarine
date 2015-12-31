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
		api := main.NewWebAPI("http://localhost:3000")
		api.Client.Transport = &webAPIMock{typhenapi.NewJSONSerializer()}

		Convey("should send a ping request", func() {
			done := make(chan *submarine.PingObject)
			go func() {
				res, _ := api.Ping("PING")
				done <- res
			}()
			res := <-done
			So(res.Message, ShouldEqual, "PING PONG")
		})
	})
}

type webAPIMock struct {
	serializer typhenapi.Serializer
}

func (mock *webAPIMock) RoundTrip(request *http.Request) (*http.Response, error) {
	response := &http.Response{Header: make(http.Header), Request: request}
	response.Header.Set("Content-Type", "application/json")
	data, _ := ioutil.ReadAll(request.Body)
	typhenType, statusCode := mock.Routes(request.URL.Path, data)

	response.StatusCode = statusCode
	body, _ := typhenType.Bytes(mock.serializer)
	response.Body = ioutil.NopCloser(bytes.NewReader(body))
	if response.StatusCode >= 400 {
		return response, errors.New("Server Error")
	}
	return response, nil
}

func (mock *webAPIMock) Routes(path string, data []byte) (typhenapi.Type, int) {
	switch path {
	case "/ping":
		params := new(webapi.PingRequestBody)
		mock.serializer.Deserialize(data, params)
		return mock.Ping(params)
	default:
		return &typhenapi.Void{}, http.StatusOK
	}
}

func (mock *webAPIMock) Ping(params *webapi.PingRequestBody) (typhenapi.Type, int) {
	typhenType := &submarine.PingObject{"PING PONG"}
	return typhenType, http.StatusOK
}
