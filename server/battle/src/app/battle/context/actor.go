package context

import (
	"app/typhenapi/type/submarine/battle"
	"github.com/chuckpreslar/emission"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	UserID() int64
	ActorType() battle.ActorType
	Event() *emission.Emitter
	Destroy()

	Position() battle.Vector
	Start()
	Update()
	OnDestroy()
}
