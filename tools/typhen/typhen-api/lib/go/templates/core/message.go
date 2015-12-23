package typhenapi

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const messageTypeLength = 4

// Message is a socket message for TyphenAPI.
type Message struct {
	Type int32
	Body []byte
}

// NewMessage creates a Message from a TyphenAPI type
func NewMessage(serializer *Serializer, messageType int32, v interface{}) (message *Message, err error) {
	typhenType := v.(Coercer)

	if typhenType == nil {
		return nil, errors.New("No TyphenAPI type")
	}

	if err := typhenType.Coerce(); err != nil {
		return nil, err
	}

	body, err := serializer.Serialize(v)
	if err != nil {
		return nil, err
	}

	return &Message{messageType, body}, nil
}

// NewMessageFromBytes creates a Message from bytes
func NewMessageFromBytes(data []byte) (message *Message, err error) {
	reader := bytes.NewReader(data)

	messageTypeBytes := make([]byte, messageTypeLength)
	if _, err := reader.Read(messageTypeBytes); err != nil {
		return nil, err
	}

	messageType := int32(binary.LittleEndian.Uint32(messageTypeBytes))

	messageBody := make([]byte, len(data)-messageTypeLength)
	if _, err := reader.Read(messageBody); err != nil {
		return nil, err
	}

	return &Message{messageType, messageBody}, nil
}

// Bytes returns bytes of the message.
func (message *Message) Bytes() []byte {
	buffer := &bytes.Buffer{}
	binary.Write(buffer, binary.LittleEndian, message.Type)
	buffer.Write(message.Body)

	return buffer.Bytes()
}
