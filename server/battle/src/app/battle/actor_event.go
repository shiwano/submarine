package battle

// ActorEvent represents a specific value for actor's event.
type ActorEvent int

const (
	actorCreated ActorEvent = iota
	actorMoved
	actorDestroyed
)
