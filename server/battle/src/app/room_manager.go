package main

import (
	"app/resp"
)

// RoomManager manages rooms.
type RoomManager struct {
	rooms           map[int64]*Room
	getOrCreateRoom chan *resp.Respondable
	deleteRoom      chan *Room
}

func newRoomManager() *RoomManager {
	roomManager := &RoomManager{
		rooms:           make(map[int64]*Room),
		getOrCreateRoom: make(chan *resp.Respondable, 32),
		deleteRoom:      make(chan *Room, 8),
	}
	go roomManager.run()
	return roomManager
}

func (m *RoomManager) run() {
	for {
		select {
		case respondable := <-m.getOrCreateRoom:
			m._getOrCreateRoom(respondable)
		case room := <-m.deleteRoom:
			m._deleteRoom(room)
		}
	}
}

func (m *RoomManager) tryGetRoom(roomID int64) (*Room, error) {
	respondable := resp.New(roomID)
	m.getOrCreateRoom <- respondable
	res, err := respondable.Receive()
	if err != nil {
		return nil, err
	}
	return res.(*Room), nil
}

func (m *RoomManager) _getOrCreateRoom(respondable *resp.Respondable) {
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
			m.deleteRoom <- room
		}
		m.rooms[roomID] = room
	}
	respondable.Respond(room, nil)
}

func (m *RoomManager) _deleteRoom(room *Room) {
	room.closeHandler = nil
	delete(m.rooms, room.id)
}
