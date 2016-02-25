package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *context.Context
	event     *event.Emitter
}

func newActor(battleContext *context.Context, userID int64, actorType battle.ActorType) *actor {
	return &actor{
		id:        battleContext.NextActorID(),
		userID:    userID,
		actorType: actorType,
		context:   battleContext,
		event:     event.New(),
	}
}

func (a *actor) ID() int64 {
	return a.id
}

func (a *actor) UserID() int64 {
	return a.userID
}

func (a *actor) ActorType() battle.ActorType {
	return a.actorType
}

func (a *actor) Event() *event.Emitter {
	return a.event
}

func (a *actor) Destroy() {
	a.context.Event.Emit(event.ActorDestroy, a)
}

// Overridable methods.
func (a *actor) Position() battle.Vector { return battle.Vector{X: 0, Y: 0} }
func (a *actor) Start()                  {}
func (a *actor) Update()                 {}
func (a *actor) OnDestroy()              {}
