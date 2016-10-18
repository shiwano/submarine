package context

import (
	"time"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// ActorParams represents actor parameters.
type ActorParams interface {
	Type() battleAPI.ActorType
	AccelMaxSpeed() float64
	AccelDuration() time.Duration
}

type actorParams struct {
	actorType     battleAPI.ActorType
	accelMaxSpeed float64
	accelDuration time.Duration
}

func (a *actorParams) Type() battleAPI.ActorType {
	return a.actorType
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
