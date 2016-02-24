package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/chuckpreslar/emission"
)

type actor struct {
	id        int64
	actorType battle.ActorType
	context   *Context
	event     *emission.Emitter
}

func newSubmarine(battleContext *Context) Actor {
	a := &actor{
		id:        battleContext.NextActorID(),
		actorType: battle.ActorType_Submarine,
		context:   battleContext,
		event:     emission.NewEmitter(),
	}
	a.context.Event.EmitSync(event.ActorCreated, a)
	return a
}

func (a *actor) ID() int64 {
	return a.id
}

func (a *actor) UserID() int64 {
	return 1
}

func (a *actor) ActorType() battle.ActorType {
	return a.actorType
}

func (a *actor) Event() *emission.Emitter {
	return a.event
}

func (a *actor) Destroy() {
	a.context.Event.EmitSync(event.ActorDestroyed, a)
}

func (a *actor) Position() battle.Vector { return battle.Vector{X: 0, Y: 0} }
func (a *actor) Start()                  {}
func (a *actor) Update()                 {}
func (a *actor) OnDestroy()              {}
