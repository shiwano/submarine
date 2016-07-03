package context

import (
	"github.com/ungerik/go3d/float64/vec2"
	"time"
)

// User represents an user in the battle.
type User struct {
	ID                            int64
	StartPosition                 *vec2.T
	SubmarineAccelerationMaxSpeed float64
	SubmarineAccelerationDuration time.Duration
	TorpedoStartOffsetLength      float64
	TorpedoAccelerationMaxSpeed   float64
	TorpedoAccelerationDuration   time.Duration
}
