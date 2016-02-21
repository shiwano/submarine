package battle

import (
	"app/typhenapi/type/submarine/battle"
	"github.com/chuckpreslar/emission"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	UserID() int64
	ActorType() battle.ActorType
	Event() *emission.Emitter
	Destroy()

	Position() battle.Vector
	Start()
	Update()
	OnDestroy()
}

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *Context
	event     *emission.Emitter
}

// ID returns the actor id.
func (a *actor) ID() int64 {
	return a.id
}

// UserID returns the user id, who has ownership of the actor.
func (a *actor) UserID() int64 {
	return a.userID
}

// ActorType returns the actor type.
func (a *actor) ActorType() battle.ActorType {
	return a.actorType
}

// Event returns the actor event.
func (a *actor) Event() *emission.Emitter {
	return a.event
}

// Destroy the actor.
func (a *actor) Destroy() {
	a.context.container.destroyActor(a)
}

// Position returns the current actor position.
func (a *actor) Position() battle.Vector {
	return battle.Vector{X: 0, Y: 0}
}

// Start the actor. This method is called when the actor is created.
func (a *actor) Start() {}

// Update the actor. This method is called every frame.
func (a *actor) Update() {}

// OnDestroy is called when the actor is destroyed.
func (a *actor) OnDestroy() {}
