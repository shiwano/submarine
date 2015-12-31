package typhenapi

import (
	"bytes"
	"errors"
	"github.com/ugorji/go/codec"
)

// Serializer serialize/deserializer a TyphenAPI type.
type Serializer interface {
	Serialize(v interface{}) (data []byte, err error)
	Deserialize(b []byte, v interface{}) error
}

// CodecSerializer is the Serializer/Deserializer for TyphenAPI types.
type CodecSerializer struct {
	Handle codec.Handle
}

// NewJSONSerializer creates an CodecSerializer for JSON format.
func NewJSONSerializer() *CodecSerializer {
	handle := &codec.JsonHandle{}
	return &CodecSerializer{handle}
}

// Serialize an TyphenAPI type to bytes.
func (s *CodecSerializer) Serialize(v interface{}) (data []byte, err error) {
	buffer := &bytes.Buffer{}

	encoder := codec.NewEncoder(buffer, s.Handle)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Deserialize bytes to an TyphenAPI type.
func (s *CodecSerializer) Deserialize(b []byte, v interface{}) error {
	reader := bytes.NewReader(b)
	decoder := codec.NewDecoder(reader, s.Handle)

	if err := decoder.Decode(v); err != nil {
		return err
	}

	typhenType := v.(Type)
	if typhenType == nil {
		return errors.New("No TyphenAPI type")
	}

	if err := typhenType.Coerce(); err != nil {
		return err
	}

	return nil
}
