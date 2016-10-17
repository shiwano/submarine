//go:generate gen

package context

import (
	"fmt"
	"time"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// Player represents a player in the battle.
// +gen * slice:"All,Any,First,Where,Count,Select[int64],GroupBy[string]"
type Player struct {
	ID              int64
	IsUser          bool
	AI              AI
	TeamLayer       navmesh.LayerMask
	StartPosition   *vec2.T
	SubmarineParams *SubmarineParams
	TorpedoParams   *TorpedoParams
}

// NewPlayer creates a player.
func NewPlayer(playerID int64, isUser bool, teamLayer navmesh.LayerMask,
	startPosition *vec2.T) *Player {
	return &Player{
		ID:            playerID,
		IsUser:        isUser,
		TeamLayer:     teamLayer,
		StartPosition: startPosition,
		SubmarineParams: &SubmarineParams{
			actorParams: &actorParams{
				accelMaxSpeed: 6,
				accelDuration: 2 * time.Second,
			},
		},
		TorpedoParams: &TorpedoParams{
			actorParams: &actorParams{
				accelMaxSpeed: 10,
				accelDuration: 1 * time.Second,
			},
			StartOffsetDistance: 1.2,
		},
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Player(%v)", p.ID)
}

// GroupByTeamLayer groups elements into a map keyed by teamLayer.
func (rcv PlayerSlice) GroupByTeamLayer(fn func(*Player) navmesh.LayerMask) map[navmesh.LayerMask]PlayerSlice {
	result := make(map[navmesh.LayerMask]PlayerSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
