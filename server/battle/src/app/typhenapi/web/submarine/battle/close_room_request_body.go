// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
	"fmt"
	"net/url"
)

var _ = errors.New

type CloseRoomRequestBody struct {
	RoomId int `codec:"room_id"`
}

// Coerce the fields.
func (t *CloseRoomRequestBody) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *CloseRoomRequestBody) Bytes(serializer *typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// QueryString returns the query string.
func (t *CloseRoomRequestBody) QueryString() string {
	queryString := fmt.Sprintf("room_id=%v", t.RoomId)
	return url.QueryEscape(queryString)
}
