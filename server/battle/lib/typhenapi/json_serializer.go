package typhenapi

import (
	"bytes"
	"encoding/json"
	"errors"
)

// JSONSerializer is JSON Serializer/Deserializer for TyphenAPI types.
type JSONSerializer struct{}

// Serialize a TyphenAPI type to bytes.
func (s *JSONSerializer) Serialize(v interface{}) (data []byte, err error) {
	b := &bytes.Buffer{}

	encoder := json.NewEncoder(b)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Deserialize bytes to a TyphenAPI type.
func (s *JSONSerializer) Deserialize(b []byte, v interface{}) error {
	r := bytes.NewReader(b)

	decoder := json.NewDecoder(r)
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
