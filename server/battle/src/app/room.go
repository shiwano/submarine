package main

// Room represents a network group for battle.
type Room struct {
	battleID uint64
	sessions map[uint64]*Session
}

func newRoom(battleID uint64) *Room {
	return &Room{battleID, make(map[uint64]*Session)}
}

func (room *Room) isEmpty() bool {
	return len(room.sessions) == 0
}

func (room *Room) join(session *Session) {
	room.sessions[session.id] = session
}

func (room *Room) leave(session *Session) {
	delete(room.sessions, session.id)
}

func (room *Room) close() {
	for _, session := range room.sessions {
		session.close()
	}
	room.sessions = nil
}
