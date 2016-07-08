package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/k0kubun/pp"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
)

var p = pp.Println

type actor struct {
	user        *context.User
	actorType   battle.ActorType
	context     *context.Context
	event       *event.Emitter
	isDestroyed bool
	motor       *motor
	stageAgent  *navmesh.Agent
}

func newActor(battleContext *context.Context, user *context.User, actorType battle.ActorType,
	position *vec2.T, direction float64, params context.ActorParams) *actor {
	a := &actor{
		user:       user,
		actorType:  actorType,
		context:    battleContext,
		event:      event.New(),
		motor:      newMotor(battleContext, position, direction, params.AccelMaxSpeed(), params.AccelDuration()),
		stageAgent: battleContext.Stage.CreateAgent(21, position),
	}

	switch actorType {
	case battle.ActorType_Submarine:
		a.stageAgent.SetLayer(context.LayerSubmarine)
	case battle.ActorType_Torpedo:
		a.stageAgent.SetLayer(context.LayerTorpedo)
	}
	a.stageAgent.SetLayer(a.User().TeamLayer)
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
	a.isDestroyed = true
	a.stageAgent.Destroy()
	a.context.Event.Emit(event.ActorDestroy, a)
}

func (a *actor) IsDestroyed() bool {
	return a.isDestroyed
}

func (a *actor) Movement() *battle.Movement {
	return a.motor.toAPIType(a.ID())
}

func (a *actor) Position() *vec2.T {
	return a.stageAgent.Position()
}

func (a *actor) BeforeUpdate() {
	position := a.motor.position()
	a.stageAgent.Move(position, 0)
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
	if a.IsDestroyed() {
		return
	}
	if obj == nil {
		a.event.Emit(event.ActorCollideWithStage, point)
	} else if other := a.context.Actor(obj.ID()); other != nil {
		a.event.Emit(event.ActorCollideWithOtherActor, other, point)
	}
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}
