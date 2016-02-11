package battle

import (
	"app/typhenapi/type/submarine/battle"
)

// Actor represents an actor in the battle.
type Actor interface {
	ID() int64
	UserID() int64
	ActorType() battle.ActorType
	Position() battle.Vector
	Update()
}

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *Context
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

// Position returns the current actor position.
func (a *actor) Position() battle.Vector {
	return battle.Vector{X: 0, Y: 0}
}

// Update the actor. This function is called every frame.
func (a *actor) Update() {}
