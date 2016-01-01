package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

// Log is a logrus.Logger
var Log = logrus.New()

// NewEngine creates a gin.Engine.
func NewEngine() (*gin.Engine, *io.PipeWriter) {
	logWriter := Log.Writer()
	roomManager := newRoomManager()

	r := gin.New()
	r.Use(gin.Recovery(), gin.LoggerWithWriter(logWriter))

	r.GET("/rooms/:id", func(c *gin.Context) {
		roomID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			Log.Error(err)
			c.String(http.StatusBadRequest, "Invalid room id.")
			return
		}

		session := newSession(1, roomID)
		if err := session.Connect(c.Writer, c.Request); err != nil {
			Log.Error(err)
			c.String(http.StatusBadRequest, "Failed to upgrade the request to web socket protocol.")
			return
		}

		roomManager.JoinToRoom <- session
		Log.Infof("Session(%v) is created", session.id)
	})

	return r, logWriter
}

func main() {
	engine, logWriter := NewEngine()
	defer logWriter.Close()
	engine.Run(":5000")
}
