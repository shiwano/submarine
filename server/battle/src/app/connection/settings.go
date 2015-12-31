package connection

import (
	"time"
)

// Settings represents a connection configuration.
type Settings struct {
	WriteWait         time.Duration
	PongWait          time.Duration
	PingPeriod        time.Duration
	MaxMessageSize    int64
	MessageBufferSize int
	ReadBufferSize    int
	WriteBufferSize   int
}

// NewSettings creates a default Settings.
func NewSettings() *Settings {
	return &Settings{
		WriteWait:         10 * time.Second,
		PongWait:          60 * time.Second,
		PingPeriod:        (60 * time.Second * 9) / 10,
		MaxMessageSize:    512,
		MessageBufferSize: 512,
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
	}
}
