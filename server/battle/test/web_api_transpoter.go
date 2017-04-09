package test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	webAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	webAPI_battle "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine/battle"
)

// WebAPITransporter is a mock object for test.
type WebAPITransporter struct {
	serializer typhenapi.Serializer
}

// NewWebAPITransporter creates a WebAPITransporter.
func NewWebAPITransporter() *WebAPITransporter {
	return &WebAPITransporter{
		serializer: new(typhenapi.MessagePackSerializer),
	}
}

// RoundTrip implements http.RoundTripper.
func (t *WebAPITransporter) RoundTrip(request *http.Request) (*http.Response, error) {
	response := &http.Response{Header: make(http.Header), Request: request}
	response.Header.Set("Content-Type", "github.com/shiwano/submarine/server/battle/application/json")
	data, _ := ioutil.ReadAll(request.Body)
	typhenType, statusCode := t.routes(request.URL.Path, data)
	err := typhenType.Coerce()
	if err != nil {
		return response, err
	}

	response.StatusCode = statusCode
	body, _ := typhenType.Bytes(t.serializer)
	response.Body = ioutil.NopCloser(bytes.NewReader(body))
	if response.StatusCode >= 400 {
		return nil, errors.New("Server Error")
	}
	return response, nil
}

func (t *WebAPITransporter) routes(path string, data []byte) (typhenapi.Type, int) {
	switch path {
	case "/ping":
		params := new(webAPI.PingRequestBody)
		t.serializer.Deserialize(data, params)
		return t.ping(params)
	case "/battle/find_room":
		params := new(webAPI_battle.FindRoomRequestBody)
		t.serializer.Deserialize(data, params)
		return t.findRoom(params)
	case "/battle/find_room_member":
		params := new(webAPI_battle.FindRoomMemberRequestBody)
		t.serializer.Deserialize(data, params)
		return t.findRoomMember(params)
	case "/battle/close_room":
		params := new(webAPI_battle.CloseRoomRequestBody)
		t.serializer.Deserialize(data, params)
		return t.closeRoom(params)
	default:
		return new(typhenapi.Void), http.StatusNotFound
	}
}

func (t *WebAPITransporter) ping(params *webAPI.PingRequestBody) (typhenapi.Type, int) {
	typhenType := &api.PingObject{Message: params.Message + " PONG"}
	return typhenType, http.StatusOK
}

func (t *WebAPITransporter) findRoomMember(params *webAPI_battle.FindRoomMemberRequestBody) (typhenapi.Type, int) {
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

func (t *WebAPITransporter) findRoom(params *webAPI_battle.FindRoomRequestBody) (typhenapi.Type, int) {
	typhenType := new(battleAPI.FindRoomObject)

	if params.RoomId == 1 {
		typhenType.Room = &battleAPI.PlayableRoom{
			Id: params.RoomId,
		}
	}
	return typhenType, http.StatusOK
}

func (t *WebAPITransporter) closeRoom(params *webAPI_battle.CloseRoomRequestBody) (typhenapi.Type, int) {
	if params.RoomId <= 0 {
		return new(typhenapi.Void), http.StatusInternalServerError
	}
	return new(typhenapi.Void), http.StatusOK
}
