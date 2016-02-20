package battle

// EventType represents a specific value for battle event types.
type EventType int

// EventType constants.
const (
	ActorCreated EventType = iota + 1
	ActorMoved
	ActorDestroyed

	AccelerationRequested
	BrakeRequested
	TurnRequested
	TorpedoRequested
	PingerRequested
)
