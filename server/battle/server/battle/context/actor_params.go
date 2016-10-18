package context

import (
	"time"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// ActorParams represents actor parameters.
type ActorParams interface {
	Type() battleAPI.ActorType
	HasLight() bool
	AccelMaxSpeed() float64
	AccelDuration() time.Duration
}

type actorParams struct {
	actorType     battleAPI.ActorType
	hasLight      bool
	accelMaxSpeed float64
	accelDuration time.Duration
}

func (a *actorParams) Type() battleAPI.ActorType {
	return a.actorType
}

func (a *actorParams) HasLight() bool {
	return a.hasLight
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
