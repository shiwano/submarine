package context

import (
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"

	"github.com/ungerik/go3d/float64/vec2"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	Player() *Player
	Type() battleAPI.ActorType
	Event() *event.Emitter

	IsDestroyed() bool
	Movement() *battleAPI.Movement
	Position() *vec2.T
	Direction() float64
	IsAccelerating() bool
	IsVisibleFrom(*Player) bool

	Destroy()

	Start()
	BeforeUpdate()
	Update()
	OnDestroy()
}
