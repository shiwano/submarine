package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

// Server represents a battle server.
type Server struct {
	*gin.Engine
	logWriter   *io.PipeWriter
	roomManager *RoomManager
}

// NewServer creates a Server.
func NewServer() *Server {
	server := &Server{gin.New(), Log.Writer(), newRoomManager()}
	server.Use(gin.Recovery(), gin.LoggerWithWriter(server.logWriter))

	server.GET("/rooms/:id", server.roomsGET)

	return server
}

func (s *Server) roomsGET(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		Log.Error(err)
		c.String(http.StatusBadRequest, "Invalid room id.")
		return
	}

	room, err := s.roomManager.tryGetRoom(roomID)
	if err != nil {
		Log.Error(err)
		c.String(http.StatusBadRequest, "Failed to get or create the room.\n")
		return
	}

	roomKey := c.Query("room_key")
	roomMember := room.findRoomMember(roomKey)
	if roomMember == nil {
		c.String(http.StatusBadRequest, "Disallow the user as the room member.")
		return
	}

	session := newSession(roomMember, roomID)
	if err := session.Connect(c.Writer, c.Request); err != nil {
		Log.Error(err)
		c.String(http.StatusBadRequest, "Failed to upgrade the request to web socket protocol.")
		return
	}

	room.join <- session
	Log.Infof("Session(%v) is created", session.id)
}
