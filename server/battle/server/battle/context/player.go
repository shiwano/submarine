//go:generate gen

package context

import (
	"fmt"
	"time"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// PlayersByTeam represents player slices grouped by team layer.
type PlayersByTeam map[navmesh.LayerMask]PlayerSlice

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
				actorType:       battleAPI.ActorType_Submarine,
				hasLight:        true,
				isAlwaysVisible: false,
				accelMaxSpeed:   6,
				accelDuration:   2 * time.Second,
			},
		},
		TorpedoParams: &TorpedoParams{
			actorParams: &actorParams{
				actorType:       battleAPI.ActorType_Torpedo,
				hasLight:        false,
				isAlwaysVisible: true,
				accelMaxSpeed:   10,
				accelDuration:   1 * time.Second,
			},
			StartOffsetDistance: 1.2,
		},
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Player(%v)", p.ID)
}

// GroupByTeam groups elements into a map keyed by team's navmesh Layer.
func (rcv PlayerSlice) GroupByTeam() PlayersByTeam {
	result := make(map[navmesh.LayerMask]PlayerSlice)
	for _, v := range rcv {
		key := v.TeamLayer
		result[key] = append(result[key], v)
	}
	return result
}
