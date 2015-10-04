package main

import (
	"app/typhen_api/core"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"io"
)

// Log is a logrus.Logger
var Log = logrus.New()

// NewEngine creates a gin.Engine.
func NewEngine() (*gin.Engine, *io.PipeWriter) {
	logWriter := Log.Writer()
	serializer := typhenapi.NewJSONSerializer()
	rooms := make(map[int]*Room)

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))
	m := melody.New()

	r.GET("/room/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		Log.Info("room id: " + roomID)
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(rawSession *melody.Session) {
		room, existsRoom := rooms[1]
		if !existsRoom {
			room = &Room{serializer, make(map[*melody.Session]*Session)}
			rooms[1] = room
		}
		room.join(rawSession)
	})

	m.HandleDisconnect(func(rawSession *melody.Session) {
		room, existsRoom := rooms[1]
		if !existsRoom {
			Log.Warn("No room exists")
			return
		}
		room.leave(rawSession)
	})

	m.HandleMessageBinary(func(rawSession *melody.Session, data []byte) {
		room, existsRoom := rooms[1]
		if !existsRoom {
			Log.Warn("No room exists")
			return
		}
		room.handleMessage(rawSession, data)
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}
