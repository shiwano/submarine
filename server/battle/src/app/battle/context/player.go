package context

import (
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"time"
)

// Player represents a player in the battle.
type Player struct {
	ID              int64
	AI              AI
	TeamLayer       navmesh.LayerMask
	StartPosition   *vec2.T
	SubmarineParams *SubmarineParams
	TorpedoParams   *TorpedoParams
}

// NewPlayer creates a player.
func NewPlayer(playerID int64, teamLayer navmesh.LayerMask, startPosition *vec2.T) *Player {
	return &Player{
		ID:            playerID,
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
