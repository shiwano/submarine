package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	UserID() int64
	ActorType() battle.ActorType
	Event() *event.Emitter
	Destroy()

	Position() battle.Vector
	Start()
	Update()
	OnDestroy()
}
