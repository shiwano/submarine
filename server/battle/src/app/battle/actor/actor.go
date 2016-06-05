package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/k0kubun/pp"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"time"
)

var p = pp.Println

const (
	accelerationMaxSpeed = 6
	accelerationDuration = 3 * time.Second
)

type actor struct {
	id         int64
	user       *context.User
	actorType  battle.ActorType
	context    *context.Context
	event      *event.Emitter
	motor      *motor
	stageAgent *navmesh.Agent
}

func newActor(battleContext *context.Context, user *context.User, actorType battle.ActorType) *actor {
	return &actor{
		id:         battleContext.NextActorID(),
		user:       user,
		actorType:  actorType,
		context:    battleContext,
		event:      event.New(),
		motor:      newMotor(battleContext, &vec2.Zero, accelerationMaxSpeed, accelerationDuration),
		stageAgent: battleContext.Stage.CreateAgent(21, &vec2.Zero),
	}
}

func (a *actor) ID() int64 {
	return a.id
}

func (a *actor) User() *context.User {
	return a.user
}

func (a *actor) Type() battle.ActorType {
	return a.actorType
}

func (a *actor) Event() *event.Emitter {
	return a.event
}

func (a *actor) Destroy() {
	a.stageAgent.Destroy()
	a.context.Event.Emit(event.ActorDestroy, a)
}

func (a *actor) Movement() *battle.Movement {
	return a.motor.toAPIType(a.id)
}

func (a *actor) Position() *vec2.T {
	return a.stageAgent.Position()
}

func (a *actor) BeforeUpdate() {
	position := a.motor.position()
	if !a.stageAgent.MoveWithValidation(position) {
		a.motor.idle(a.stageAgent.Position())
		a.context.Event.Emit(event.ActorMove, a)
	}
}

func (a *actor) accelerate(direction float64) {
	a.motor.accelerate()
	a.motor.turn(direction)
	a.context.Event.Emit(event.ActorMove, a)
}

func (a *actor) brake(direction float64) {
	a.motor.brake()
	a.motor.turn(direction)
	a.context.Event.Emit(event.ActorMove, a)
}

func (a *actor) turn(direction float64) {
	a.motor.turn(direction)
	a.context.Event.Emit(event.ActorMove, a)
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}
