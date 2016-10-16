package server

import (
	"github.com/shiwano/submarine/server/battle/lib/respondable"
)

type roomManager struct {
	rooms             map[int64]*room
	getOrCreateRoomCh chan *respondable.T
	deleteRoomCh      chan *room
}

func newRoomManager() *roomManager {
	rm := &roomManager{
		rooms:             make(map[int64]*room),
		getOrCreateRoomCh: make(chan *respondable.T, 32),
		deleteRoomCh:      make(chan *room, 8),
	}
	go rm.run()
	return rm
}

func (rm *roomManager) run() {
	for {
		select {
		case respondable := <-rm.getOrCreateRoomCh:
			rm.getOrCreateRoom(respondable)
		case room := <-rm.deleteRoomCh:
			rm.deleteRoom(room)
		}
	}
}

func (rm *roomManager) fetchRoom(roomID int64) (*room, error) {
	respondable := respondable.New(roomID)
	rm.getOrCreateRoomCh <- respondable
	res, err := respondable.Receive()
	if err != nil {
		return nil, err
	}
	return res.(*room), nil
}

func (rm *roomManager) getOrCreateRoom(respondable *respondable.T) {
	roomID := respondable.Value.(int64)
	r, ok := rm.rooms[roomID]
	if !ok {
		newRoom, err := newRoom(roomID)
		if err != nil {
			respondable.Respond(nil, err)
			return
		}

		r = newRoom
		r.closeHandler = func(r *room) {
			rm.deleteRoomCh <- r
		}
		rm.rooms[roomID] = r
	}
	respondable.Respond(r, nil)
}

func (rm *roomManager) deleteRoom(r *room) {
	r.closeHandler = nil
	delete(rm.rooms, r.id)
}
