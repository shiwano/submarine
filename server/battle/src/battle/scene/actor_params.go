package scene

import (
	"time"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// ActorParams represents actor parameters.
type ActorParams interface {
	Type() battleAPI.ActorType
	HasLight() bool
	IsAlwaysVisible() bool
	AccelMaxSpeed() float64
	AccelDuration() time.Duration
}

type actorParams struct {
	actorType       battleAPI.ActorType
	hasLight        bool
	isAlwaysVisible bool
	accelMaxSpeed   float64
	accelDuration   time.Duration
}

func (a *actorParams) Type() battleAPI.ActorType    { return a.actorType }
func (a *actorParams) HasLight() bool               { return a.hasLight }
func (a *actorParams) IsAlwaysVisible() bool        { return a.isAlwaysVisible }
func (a *actorParams) AccelMaxSpeed() float64       { return a.accelMaxSpeed }
func (a *actorParams) AccelDuration() time.Duration { return a.accelDuration }

// SubmarineParams represents submarine parameters.
type SubmarineParams struct {
	*actorParams
	TorpedoCount           int64
	TorpedoCooldownSeconds float64
	PingerIntervalSeconds  float64
	PingerCooldownSeconds  float64
	WatcherCooldownSeconds float64
}

// TorpedoParams represents torpedo parameters.
type TorpedoParams struct {
	*actorParams
	StartOffsetDistance float64
}

// WatcherParams represents watcher parameters.
type WatcherParams struct {
	*actorParams
	UptimeSeconds float64
}
