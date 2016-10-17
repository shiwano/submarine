//go:generate gen

package context

import (
	"github.com/ungerik/go3d/float64/vec2"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/event"
)

// Actor represents an actor in the battle.
// +gen slice:"All,Any,First,MaxBy,MinBy,SortBy,Where,Count,GroupBy[string],GroupBy[*Player]"
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
