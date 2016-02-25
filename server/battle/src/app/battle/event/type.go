package event

// Type represents a specific value for battle event types.
type Type int

// Type constants.
const (
	ActorCreated Type = iota + 1
	ActorDestroyed

	ActorAdded
	ActorMoved
	ActorRemoved

	AccelerationRequested
	BrakeRequested
	TurnRequested
	TorpedoRequested
	PingerRequested
)
