package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	webapi "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	"github.com/shiwano/submarine/server/battle/src/config"
	"github.com/shiwano/submarine/server/battle/src/logger"
	"github.com/shiwano/submarine/server/battle/src/room"
)

// Server represents a battle server.
type Server struct {
	*gin.Engine
	logWriter   *io.PipeWriter
	roomManager *room.Manager
	webAPI      *webapi.WebAPI
}

// New creates a Server.
func New() *Server {
	ctx := context.Background()
	webAPI := newWebAPI(config.Config.ApiServerBaseUri)
	server := &Server{
		Engine:      gin.New(),
		logWriter:   logger.Log.Writer(),
		roomManager: room.NewManager(ctx, webAPI),
		webAPI:      webAPI,
	}
	server.Use(gin.Recovery(), gin.LoggerWithWriter(server.logWriter))

	server.GET("/rooms/:id", server.roomsGET)

	return server
}

func (s *Server) roomsGET(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || roomID <= 0 {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Invalid room id")
		return
	}

	res, err := s.webAPI.Battle.FindRoomMember(c.Query("room_key"))
	if err != nil {
		logger.Log.Error(err)
		c.String(http.StatusInternalServerError, "Failed to authenticate the room key")
		return
	}
	if res.RoomMember == nil {
		c.String(http.StatusForbidden, "Invalid room key")
		return
	}

	room, err := s.roomManager.FetchRoom(roomID)
	if err != nil {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Failed to fetch the room")
		return
	}

	if err := room.Join(res.RoomMember, c.Writer, c.Request); err != nil {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Failed to upgrade the connection to Web Socket Protocol")
		return
	}
}

func init() {
	switch config.Env {
	case "test":
		gin.SetMode(gin.TestMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	}
}
