package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/k0kubun/pp"
	"github.com/ungerik/go3d/float64/vec2"
	"time"
)

var p = pp.Println

const (
	accelerationMaxSpeed = 2
	accelerationDuration = 3 * time.Second
)

type actor struct {
	id        int64
	userID    int64
	actorType battle.ActorType
	context   *context.Context
	event     *event.Emitter
	motor     *motor
}

func newActor(battleContext *context.Context, userID int64, actorType battle.ActorType) *actor {
	return &actor{
		id:        battleContext.NextActorID(),
		userID:    userID,
		actorType: actorType,
		context:   battleContext,
		event:     event.New(),
		motor: newMotor(battleContext, &vec2.Zero,
			accelerationMaxSpeed, accelerationDuration),
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
	return a.motor.toAPIType(a.id)
}

func (a *actor) Position() *vec2.T {
	return a.motor.position()
}

func (a *actor) accelerate() {
	a.motor.accelerate()
	a.event.Emit(event.ActorMove, a)
}

func (a *actor) brake() {
	a.motor.brake()
	a.event.Emit(event.ActorMove, a)
}

func (a *actor) turn(direction float64) {
	a.motor.turn(direction)
	a.event.Emit(event.ActorMove, a)
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}
