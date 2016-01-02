package connection

import (
	"time"
)

// Settings represents connection settings.
type Settings struct {
	WriteWait                time.Duration
	PongWait                 time.Duration
	PingPeriod               time.Duration
	MessageChannelBufferSize int
	MaxMessageSize           int64
	ReadBufferSize           int
	WriteBufferSize          int
}

// newDefaultSettings creates default settings.
func newDefaultSettings() *Settings {
	return &Settings{
		WriteWait:                10 * time.Second,
		PongWait:                 60 * time.Second,
		PingPeriod:               (60 * time.Second * 9) / 10,
		MessageChannelBufferSize: 256,
		MaxMessageSize:           2048,
		ReadBufferSize:           4096,
		WriteBufferSize:          4096,
	}
}
