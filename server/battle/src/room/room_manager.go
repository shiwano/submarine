package room

import (
	"context"

	"github.com/shiwano/submarine/server/battle/lib/respondable"
	webAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	"github.com/shiwano/submarine/server/battle/src/logger"
)

// Manager manages rooms with goroutine safe.
type Manager struct {
	ctx                context.Context
	webAPI             *webAPI.WebAPI
	rooms              map[int64]*Room
	fetchRoomRequested chan *respondable.T
	roomClosed         chan int64
	closed             chan struct{}
}

// NewManager creates a Manager.
func NewManager(ctx context.Context, webAPI *webAPI.WebAPI) *Manager {
	rm := &Manager{
		ctx:                ctx,
		webAPI:             webAPI,
		rooms:              make(map[int64]*Room),
		fetchRoomRequested: make(chan *respondable.T),
		roomClosed:         make(chan int64),
		closed:             make(chan struct{}),
	}
	go rm.run()
	return rm
}

// Closed returns a channel that receives a value when the room manager closed.
func (rm *Manager) Closed() <-chan struct{} {
	return rm.closed
}

// FetchRoom fetches a room.
func (rm *Manager) FetchRoom(roomID int64) (*Room, error) {
	res := respondable.New(roomID)
	rm.fetchRoomRequested <- res
	v, err := res.Receive()
	return v.(*Room), err
}

func (rm *Manager) run() {
	logger.Log.Info("RoomManager opened")

loop:
	for {
		select {
		case <-rm.ctx.Done():
			break loop
		case roomID := <-rm.roomClosed:
			rm.deleteRoom(roomID)
		case res := <-rm.fetchRoomRequested:
			r, err := rm.getOrCreateRoom(res.Value.(int64))
			res.Respond(r, err)
		}
	}
	for _, r := range rm.rooms {
		<-r.closed
	}
	close(rm.closed)
	logger.Log.Info("RoomManager closed")
}

func (rm *Manager) getOrCreateRoom(roomID int64) (*Room, error) {
	if r, ok := rm.rooms[roomID]; ok {
		return r, nil
	}
	r, err := newRoom(rm.ctx, rm.webAPI, roomID)
	if err != nil {
		return nil, err
	}
	rm.rooms[roomID] = r
	go func() {
		<-r.closed

		select {
		case <-rm.ctx.Done():
		case rm.roomClosed <- roomID:
		}
	}()
	return r, nil
}

func (rm *Manager) deleteRoom(roomID int64) {
	delete(rm.rooms, roomID)
}
