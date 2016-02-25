package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/chuckpreslar/emission"
)

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *Context
	event     *emission.Emitter

	isCalledStart     bool
	isCalledUpdate    bool
	isCalledOnDestroy bool
}

func newSubmarine(battleContext *Context) *actor {
	id := battleContext.NextActorID()
	a := &actor{
		id:        id,
		userID:    id * 100,
		actorType: battle.ActorType_Submarine,
		context:   battleContext,
		event:     emission.NewEmitter(),
	}
	a.context.Event.EmitSync(event.ActorCreate, a)
	return a
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

func (a *actor) Event() *emission.Emitter {
	return a.event
}

func (a *actor) Destroy() {
	a.context.Event.EmitSync(event.ActorDestroy, a)
}

func (a *actor) Position() battle.Vector {
	return battle.Vector{X: 0, Y: 0}
}

func (a *actor) Start() {
	a.isCalledStart = true
}

func (a *actor) Update() {
	a.isCalledUpdate = true
}

func (a *actor) OnDestroy() {
	a.isCalledOnDestroy = true
}
