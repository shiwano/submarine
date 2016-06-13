package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

var lastCreateActorID int64

type actor struct {
	id        int64
	user      *User
	actorType battle.ActorType
	context   *Context
	event     *event.Emitter

	isCalledStart     bool
	isCalledUpdate    bool
	isCalledOnDestroy bool
}

func newSubmarine(battleContext *Context) *actor {
	lastCreateActorID++
	id := lastCreateActorID
	user := &User{ID: id * 100, StartPosition: &vec2.Zero}
	a := &actor{
		id:        id,
		user:      user,
		actorType: battle.ActorType_Submarine,
		context:   battleContext,
		event:     event.New(),
	}
	a.context.Event.Emit(event.ActorCreate, a)
	return a
}

func (a *actor) ID() int64 {
	return a.id
}

func (a *actor) User() *User {
	return a.user
}

func (a *actor) Type() battle.ActorType {
	return a.actorType
}

func (a *actor) Event() *event.Emitter {
	return a.event
}

func (a *actor) Destroy() {
	a.context.Event.Emit(event.ActorDestroy, a)
}

func (a *actor) Movement() *battle.Movement {
	panic("not implemented yet.")
}

func (a *actor) Position() *vec2.T {
	return &vec2.Zero
}

func (a *actor) BeforeUpdate() {
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
