package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io"
)

// Log is a logrus.Logger
var Log = logrus.New()

// NewEngine creates a gin.Engine.
func NewEngine() (*gin.Engine, *io.PipeWriter) {
	logWriter := Log.Writer()
	rooms := make(map[uint64]*Room)

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))

	r.GET("/battle", func(c *gin.Context) {
		session := newSession()

		battleID, _ := getBattleID(c.Request)
		room, existsRoom := rooms[battleID]
		if !existsRoom {
			room = newRoom(battleID)
			rooms[battleID] = room
		}
		room.join(session)
		session.room = room

		err := session.Connect(c.Writer, c.Request)
		if err != nil {
			Log.Infof("Session(%v) is created", session.id)
		}
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}
