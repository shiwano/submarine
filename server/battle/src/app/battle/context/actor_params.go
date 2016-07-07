package context

import (
	"time"
)

// ActorParams represents actor parameters.
type ActorParams interface {
	AccelMaxSpeed() float64
	AccelDuration() time.Duration
}

type actorParams struct {
	accelMaxSpeed float64
	accelDuration time.Duration
}

func (a *actorParams) AccelMaxSpeed() float64 {
	return a.accelMaxSpeed
}

func (a *actorParams) AccelDuration() time.Duration {
	return a.accelDuration
}

// SubmarineParams represents submatine parameters.
type SubmarineParams struct {
	*actorParams
}

// TorpedoParams represents submatine parameters.
type TorpedoParams struct {
	*actorParams
	StartOffsetDistance float64
}
