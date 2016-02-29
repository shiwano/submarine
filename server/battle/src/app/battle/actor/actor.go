package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/currentmillis"
	"app/typhenapi/type/submarine/battle"
	"github.com/k0kubun/pp"
	"github.com/ungerik/go3d/float64/vec2"
	"time"
)

var p = pp.Println

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *context.Context
	event     *event.Emitter

	direction float64
	movedAt   time.Time
}

func newActor(battleContext *context.Context, userID int64, actorType battle.ActorType) *actor {
	return &actor{
		id:        battleContext.NextActorID(),
		userID:    userID,
		actorType: actorType,
		context:   battleContext,
		event:     event.New(),
		movedAt:   time.Now(),
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

func (a *actor) Movement() *battle.Movement {
	position := a.Position()
	return &battle.Movement{
		ActorId:     a.id,
		Position:    &battle.Point{X: position[0], Y: position[1]},
		Direction:   a.direction,
		MovedAt:     currentmillis.Milliseconds(a.movedAt),
		Accelerator: nil,
	}
}

// Overridable methods.
func (a *actor) Position() *vec2.T { return &vec2.Zero }
func (a *actor) Start()            {}
func (a *actor) Update()           {}
func (a *actor) OnDestroy()        {}
