package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	User() *User
	Type() battle.ActorType
	Event() *event.Emitter

	IsDestroyed() bool
	Movement() *battle.Movement
	Position() *vec2.T
	Direction() float64
	IsAccelerating() bool

	Destroy()

	Start()
	BeforeUpdate()
	Update()
	OnDestroy()
}
