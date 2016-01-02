package main

// RoomManager manages rooms.
type RoomManager struct {
	rooms      map[uint64]*Room
	joinToRoom chan *Session
	deleteRoom chan *Room
	close      chan struct{}
}

func newRoomManager() *RoomManager {
	roomManager := &RoomManager{
		rooms:      make(map[uint64]*Room),
		joinToRoom: make(chan *Session, 32),
		deleteRoom: make(chan *Room, 8),
		close:      make(chan struct{}),
	}
	go roomManager.run()
	return roomManager
}

func (m *RoomManager) run() {
loop:
	for {
		select {
		case session := <-m.joinToRoom:
			m._joinToRoom(session)
		case room := <-m.deleteRoom:
			m._deleteRoom(room)
		case <-m.close:
			m._close()
			break loop
		}
	}
}

func (m *RoomManager) _joinToRoom(session *Session) {
	room, ok := m.rooms[session.roomID]
	if !ok {
		room = newRoom(session.roomID)
		room.closeHandler = func(room *Room) {
			m.deleteRoom <- room
		}
		m.rooms[session.roomID] = room
	}
	room.join <- session
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
