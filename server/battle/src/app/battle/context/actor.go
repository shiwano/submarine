package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	UserID() int64
	ActorType() battle.ActorType
	Event() *event.Emitter
	Destroy()

	Position() *vec2.T
	Start()
	Update()
	OnDestroy()
}
