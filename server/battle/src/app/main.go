package main

import (
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
	rooms := make(map[int]*Room)
	sessions := make(map[*melody.Session]*Session)

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))
	m := melody.New()

	r.GET("/room/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		Log.Info("room id: " + roomID)
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(rawSession *melody.Session) {
		session := newSession(rawSession, 1)
		sessions[rawSession] = session

		room, existsRoom := rooms[1]
		if !existsRoom {
			room = &Room{make(map[uint64]*Session)}
			rooms[1] = room
		}
		room.join(session)
		session.room = room
	})

	m.HandleDisconnect(func(rawSession *melody.Session) {
		if session, ok := sessions[rawSession]; ok && session.room != nil {
			session.room.leave(session)
			session.room = nil
		}
	})

	m.HandleMessageBinary(func(rawSession *melody.Session, data []byte) {
		if session, ok := sessions[rawSession]; ok {
			session.handleMessage(data)
		}
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}
