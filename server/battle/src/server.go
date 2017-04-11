package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	webapi "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	"github.com/shiwano/submarine/server/battle/src/config"
	"github.com/shiwano/submarine/server/battle/src/logger"
	"github.com/shiwano/submarine/server/battle/src/room"
)

// Server represents a battle server.
type Server struct {
	*http.Server
	router      *mux.Router
	roomManager *room.Manager
	webAPI      *webapi.WebAPI
}

// New creates a Server.
func New(addr string) *Server {
	ctx := context.Background()
	router := mux.NewRouter()
	webAPI := newWebAPI(config.Config.ApiServerBaseUri)
	roomManager := room.NewManager(ctx, webAPI)

	s := &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router:      router,
		roomManager: roomManager,
		webAPI:      webAPI,
	}
	s.router.HandleFunc("/rooms/{id}", s.roomsGET)
	return s
}

func (s *Server) roomsGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil || roomID <= 0 {
		s.writeString(w, http.StatusForbidden, "Invalid room id")
		return
	}

	q := r.URL.Query()
	res, err := s.webAPI.Battle.FindRoomMember(q.Get("room_key"))
	if err != nil {
		logger.Log.Error(err)
		s.writeString(w, http.StatusInternalServerError, "Failed to authenticate the room key")
		return
	}
	if res.RoomMember == nil {
		s.writeString(w, http.StatusForbidden, "Invalid room key")
		return
	}

	room, err := s.roomManager.FetchRoom(roomID)
	if err != nil {
		logger.Log.Error(err)
		s.writeString(w, http.StatusForbidden, "Failed to fetch the room")
		return
	}

	if ok := room.Join(res.RoomMember, w, r); !ok {
		s.writeString(w, http.StatusForbidden, "Failed to join into the room")
		return
	}
}

func (s *Server) writeString(w http.ResponseWriter, statusCode int, text string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(text))
	if err != nil {
		logger.Log.Error(err)
	}
}
