package main

import (
	"app/typhen_api/core"
	messagedispatcher "app/typhen_api/messagedispatcher/submarine"
	"github.com/olahol/melody"
)

// Session represents a network session that has user infos.
type Session struct {
	*melody.Session
	serializer *typhenapi.Serializer
	Dispatcher *messagedispatcher.MessageDispatcher
}

func newSession(melodySession *melody.Session, serializer *typhenapi.Serializer) *Session {
	messageDispatcher := messagedispatcher.New(serializer, func(m []byte, err error) { Log.Error(err) })
	return &Session{melodySession, serializer, messageDispatcher}
}

func (session *Session) send(messageBody interface{}) {
	message, err := typhenapi.NewMessage(session.serializer, messageBody)
	if err != nil {
		Log.Error(err)
		return
	}
	session.WriteBinary(message.Bytes())
}

func (session *Session) handleMessage(data []byte) {
	session.Dispatcher.HandleMessage(data)
}
