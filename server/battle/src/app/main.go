package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"io"
	"net/http"
	"strconv"
)

// Log is a logrus.Logger
var Log = logrus.New()

// NewEngine creates a gin.Engine.
func NewEngine() (*gin.Engine, *io.PipeWriter) {
	logWriter := Log.Writer()
	rooms := make(map[uint64]*Room)
	sessions := make(map[*melody.Session]*Session)

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))
	m := melody.New()

	r.GET("/battle", func(c *gin.Context) {
		if _, err := getBattleID(c.Request); err == nil {
			m.HandleRequest(c.Writer, c.Request)
		}
	})

	m.HandleConnect(func(rawSession *melody.Session) {
		session := newSession(rawSession, 1)
		sessions[rawSession] = session

		battleID, _ := getBattleID(rawSession.Request)
		room, existsRoom := rooms[battleID]
		if !existsRoom {
			room = newRoom(battleID)
			rooms[battleID] = room
		}
		room.join(session)
		session.room = room
	})

	m.HandleDisconnect(func(rawSession *melody.Session) {
		if session, ok := sessions[rawSession]; ok {
			delete(sessions, rawSession)

			if session.room != nil {
				session.room.leave(session)
				if session.room.isEmpty() {
					delete(rooms, session.room.battleID)
				}
				session.room = nil
			}
		}
	})

	m.HandleMessageBinary(func(rawSession *melody.Session, data []byte) {
		if session, ok := sessions[rawSession]; ok {
			session.onMessage(data)
		}
	})

	m.HandleError(func(rawSession *melody.Session, err error) {
		if session, ok := sessions[rawSession]; ok {
			session.onError(nil, err)
		}
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}

func getBattleID(request *http.Request) (uint64, error) {
	return strconv.ParseUint(request.URL.Query().Get("battle_id"), 10, 64)
}
