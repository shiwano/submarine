package main

// Room represents a network group for battle.
type Room struct {
	sessions map[uint64]*Session
}

func (room *Room) join(session *Session) {
	room.sessions[session.id] = session
}

func (room *Room) leave(session *Session) {
	delete(room.sessions, session.id)
}

func (room *Room) close() {
	for _, session := range room.sessions {
		session.Close()
	}
	room.sessions = nil
}
