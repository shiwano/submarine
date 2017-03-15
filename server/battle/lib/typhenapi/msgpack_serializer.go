package typhenapi

import (
	"bytes"
	"errors"

	"gopkg.in/vmihailenco/msgpack.v2"
)

// MessagePackSerializer is MessagePack Serializer/Deserializer for TyphenAPI types.
type MessagePackSerializer struct{}

// Serialize a TyphenAPI type to bytes.
func (s *MessagePackSerializer) Serialize(v interface{}) (data []byte, err error) {
	b := &bytes.Buffer{}

	encoder := msgpack.NewEncoder(b)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize bytes to a TyphenAPI type.
func (s *MessagePackSerializer) Deserialize(b []byte, v interface{}) error {
	r := bytes.NewReader(b)

	decoder := msgpack.NewDecoder(r)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	result := v.(Type)
	if result == nil {
		return errors.New("Not TyphenAPI type")
	}

	if err := result.Coerce(); err != nil {
		return err
	}
	return nil
}
