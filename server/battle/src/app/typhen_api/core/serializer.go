package typhenapi

import (
	"bytes"
	"errors"
	"github.com/ugorji/go/codec"
)

// Serializer is the Serializer/Deserializer for TyphenAPI types.
type Serializer struct {
	Handle codec.Handle
}

// NewJSONSerializer creates an Serializer for JSON format.
func NewJSONSerializer() *Serializer {
	handle := &codec.JsonHandle{}
	return &Serializer{handle}
}

// Serialize an TyphenAPI type to bytes.
func (s *Serializer) Serialize(v interface{}) (data []byte, err error) {
	buffer := &bytes.Buffer{}
	encoder := codec.NewEncoder(buffer, s.Handle)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Deserialize bytes to an TyphenAPI type.
func (s *Serializer) Deserialize(b []byte, v interface{}) error {
	reader := bytes.NewReader(b)
	decoder := codec.NewDecoder(reader, s.Handle)

	if err := decoder.Decode(v); err != nil {
		return err
	}

	coercer := v.(Coercer)

	if coercer == nil {
		return errors.New("No TyphenAPI type")
	}

	if err := coercer.Coerce(); err != nil {
		return err
	}

	return nil
}
