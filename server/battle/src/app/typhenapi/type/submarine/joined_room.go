// This file was generated by typhen-api

package submarine

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// JoinedRoom is a kind of TyphenAPI type.
type JoinedRoom struct {
	BattleServerBaseUri string `codec:"battle_server_base_uri"`
	RoomKey             string `codec:"room_key"`
	Id                  int    `codec:"id"`
	Members             []User `codec:"members"`
}

// Coerce the fields.
func (t *JoinedRoom) Coerce() error {
	if t.Members == nil {
		return errors.New("Members should not be empty")
	}
	return nil
}

// Bytes creates the byte array.
func (t *JoinedRoom) Bytes(serializer *typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
