package main_test

import (
	"app/typhen_api/core"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct {
	Message string `codec:"message"`
}

func (t *testType) RealTimeMessageType() int32 {
	return 1
}

func (t *testType) Coerce() error {
	if t.Message == "" {
		return errors.New("Message is empty")
	}

	return nil
}

func TestNewMessage(t *testing.T) {
	serializer := typhenapi.NewJSONSerializer()
	message, err := typhenapi.NewMessage(serializer, &testType{"Foobar"})

	assert.NoError(t, err)
	assert.EqualValues(t, 1, message.Type)

	deserialized := &testType{}
	assert.NoError(t, serializer.Deserialize(message.Body, deserialized))
	assert.Equal(t, "Foobar", deserialized.Message)
}

func TestNewMessageFromBytes(t *testing.T) {
	data := []byte{0xF0, 0xFF, 0x00, 0x00, 2, 3, 5, 7, 11}
	message, err := typhenapi.NewMessageFromBytes(data)

	assert.NoError(t, err)
	assert.EqualValues(t, 65520, message.Type)
	assert.Equal(t, []byte{2, 3, 5, 7, 11}, message.Body)
}

func TestMessageBytes(t *testing.T) {
	serializer := typhenapi.NewJSONSerializer()
	messageA, _ := typhenapi.NewMessage(serializer, &testType{"Foobar"})
	message, _ := typhenapi.NewMessageFromBytes(messageA.Bytes())

	deserialized := &testType{}
	assert.NoError(t, serializer.Deserialize(message.Body, deserialized))
	assert.Equal(t, message.Type, messageA.Type)
	assert.Equal(t, "Foobar", deserialized.Message)
}
