package main

import (
	"lib/respondable"
)

// RoomManager manages rooms.
type RoomManager struct {
	rooms             map[int64]*Room
	getOrCreateRoomCh chan *respondable.T
	deleteRoomCh      chan *Room
}

func newRoomManager() *RoomManager {
	roomManager := &RoomManager{
		rooms:             make(map[int64]*Room),
		getOrCreateRoomCh: make(chan *respondable.T, 32),
		deleteRoomCh:      make(chan *Room, 8),
	}
	go roomManager.run()
	return roomManager
}

func (m *RoomManager) run() {
	for {
		select {
		case respondable := <-m.getOrCreateRoomCh:
			m.getOrCreateRoom(respondable)
		case room := <-m.deleteRoomCh:
			m.deleteRoom(room)
		}
	}
}

func (m *RoomManager) fetchRoom(roomID int64) (*Room, error) {
	respondable := respondable.New(roomID)
	m.getOrCreateRoomCh <- respondable
	res, err := respondable.Receive()
	if err != nil {
		return nil, err
	}
	return res.(*Room), nil
}

func (m *RoomManager) getOrCreateRoom(respondable *respondable.T) {
	roomID := respondable.Value.(int64)
	room, ok := m.rooms[roomID]
	if !ok {
		newRoom, err := newRoom(roomID)
		if err != nil {
			respondable.Respond(nil, err)
			return
		}

		room = newRoom
		room.closeHandler = func(room *Room) {
			m.deleteRoomCh <- room
		}
		m.rooms[roomID] = room
	}
	respondable.Respond(room, nil)
}

func (m *RoomManager) deleteRoom(room *Room) {
	room.closeHandler = nil
	delete(m.rooms, room.id)
}
