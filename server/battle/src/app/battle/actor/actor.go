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
	user       *context.User
	actorType  battle.ActorType
	context    *context.Context
	event      *event.Emitter
	motor      *motor
	stageAgent *navmesh.Agent
}

func newActor(battleContext *context.Context, user *context.User, actorType battle.ActorType, startPos *vec2.T) *actor {
	a := &actor{
		user:       user,
		actorType:  actorType,
		context:    battleContext,
		event:      event.New(),
		motor:      newMotor(battleContext, startPos, accelerationMaxSpeed, accelerationDuration),
		stageAgent: battleContext.Stage.CreateAgent(21, startPos),
	}
	a.stageAgent.SetCollideHandler(a.onStageAgentCollide)
	return a
}

func (a *actor) ID() int64 {
	return a.stageAgent.ID()
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
	return a.motor.toAPIType(a.ID())
}

func (a *actor) Position() *vec2.T {
	return a.stageAgent.Position()
}

func (a *actor) BeforeUpdate() {
	position := a.motor.position()
	a.stageAgent.Move(position)
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

func (a *actor) idle() {
	a.motor.idle(a.stageAgent.Position())
	a.context.Event.Emit(event.ActorMove, a)
}

func (a *actor) onStageAgentCollide(obj navmesh.Object, point vec2.T) {
	if obj == nil {
		a.event.Emit(event.ActorCollide, nil, point)
	} else {
		a.event.Emit(event.ActorCollide, a.context.Actor(obj.ID()), point)
	}
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}
