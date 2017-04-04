//go:generate gen

package scene

import (
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// Actor represents an actor in the battle.
// +gen slice:"All,Any,First,MaxBy,MinBy,SortBy,Where,Count,GroupBy[string],GroupBy[*Player]"
type Actor interface {
	ID() int64
	Player() *Player
	Type() battleAPI.ActorType
	Event() *ActorEventEmitter
	IsDestroyed() bool
	Movement() *battleAPI.Movement
	Position() *vec2.T
	Direction() float64
	IsAccelerating() bool
	IsVisibleFrom(navmesh.LayerMask) bool
	Submarine(navmesh.LayerMask) *battleAPI.ActorSubmarineObject

	Start()
	BeforeUpdate()
	Update()
	AfterUpdate()
	OnDestroy()

	Destroy()
}

// GroupByLayer groups elements into a map keyed by navmesh layer.
func (rcv ActorSlice) GroupByLayer(fn func(Actor) navmesh.LayerMask) map[navmesh.LayerMask]ActorSlice {
	result := make(map[navmesh.LayerMask]ActorSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
