package main

// RoomManager manages rooms.
type RoomManager struct {
	rooms           map[uint64]*Room
	getOrCreateRoom chan *Respondable
	deleteRoom      chan *Room
	close           chan struct{}
}

func newRoomManager() *RoomManager {
	roomManager := &RoomManager{
		rooms:           make(map[uint64]*Room),
		getOrCreateRoom: make(chan *Respondable, 32),
		deleteRoom:      make(chan *Room, 8),
		close:           make(chan struct{}),
	}
	go roomManager.run()
	return roomManager
}

func (m *RoomManager) run() {
loop:
	for {
		select {
		case respondable := <-m.getOrCreateRoom:
			m._getOrCreateRoom(respondable)
		case room := <-m.deleteRoom:
			m._deleteRoom(room)
		case <-m.close:
			m._close()
			break loop
		}
	}
}

func (m *RoomManager) tryGetRoom(roomID uint64) (*Room, error) {
	respondable := newRespondable(roomID)
	m.getOrCreateRoom <- respondable
	response, err := respondable.wait()
	return response.(*Room), err
}

func (m *RoomManager) _getOrCreateRoom(respondable *Respondable) {
	roomID := respondable.value.(uint64)
	room, ok := m.rooms[roomID]
	if !ok {
		room = newRoom(roomID)
		room.closeHandler = func(room *Room) {
			m.deleteRoom <- room
		}
		m.rooms[roomID] = room
	}
	respondable.done <- room
}

func (m *RoomManager) _deleteRoom(room *Room) {
	room.closeHandler = nil
	delete(m.rooms, room.id)
}

func (m *RoomManager) _close() {
	for _, room := range m.rooms {
		m._deleteRoom(room)
		room.close <- struct{}{}
	}
}
