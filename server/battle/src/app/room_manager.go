package main

// RoomManager manages rooms.
type RoomManager struct {
	rooms      map[uint64]*Room
	JoinToRoom chan *Session
	Close      chan struct{}
	RemoveRoom chan *Room
}

func newRoomManager() *RoomManager {
	manager := &RoomManager{
		rooms:      make(map[uint64]*Room),
		JoinToRoom: make(chan *Session),
		Close:      make(chan struct{}),
	}
	go manager.run()
	return manager
}

func (r *RoomManager) run() {
loop:
	for {
		select {
		case session := <-r.JoinToRoom:
			r.joinToRoom(session)
		case room := <-r.RemoveRoom:
			r.removeRoom(room)
			break loop
		case <-r.Close:
			r.closeAllRooms()
			break loop
		}
	}
}

func (r *RoomManager) joinToRoom(session *Session) {
	room, ok := r.rooms[session.roomID]
	if !ok {
		room = newRoom(session.roomID)
		room.closeHandler = func(room *Room) {
			r.RemoveRoom <- room
		}
		r.rooms[session.roomID] = room
	}
	room.Join <- session
}

func (r *RoomManager) removeRoom(room *Room) {
	delete(r.rooms, room.id)
}

func (r *RoomManager) closeAllRooms() {
	for _, room := range r.rooms {
		room.Close <- struct{}{}
	}
}
