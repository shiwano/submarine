package context

import (
	"github.com/ungerik/go3d/float64/vec2"
	"time"
)

// User represents an user in the battle.
type User struct {
	ID              int64
	StartPosition   *vec2.T
	SubmarineParams *SubmarineParams
	TorpedoParams   *TorpedoParams
}

// NewUser creates an user.
func NewUser(userID int64, startPosition *vec2.T) *User {
	return &User{
		ID:            userID,
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
