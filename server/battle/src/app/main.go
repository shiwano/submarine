package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

// Log is a logrus.Logger
var Log = logrus.New()

// NewEngine creates a gin.Engine.
func NewEngine() (*gin.Engine, *io.PipeWriter) {
	logWriter := Log.Writer()
	rooms := make(map[uint64]*Room)

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))

	r.GET("/rooms/:id", func(c *gin.Context) {
		session := newSession()

		roomID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			Log.Error(err)
		}

		room, ok := rooms[roomID]
		if !ok {
			room = newRoom(roomID)
			rooms[roomID] = room
		}
		room.join(session)
		session.room = room

		if err := session.Connect(c.Writer, c.Request); err != nil {
			Log.Error(err)
		}
		Log.Infof("Session(%v) is created", session.id)
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}
