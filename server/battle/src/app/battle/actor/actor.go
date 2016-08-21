package actor

import (
	"app/battle/context"
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
)

type actor struct {
	player       *context.Player
	actorType    battleAPI.ActorType
	ctx          *context.Context
	event        *event.Emitter
	isDestroyed  bool
	motor        *motor
	stageAgent   *navmesh.Agent
	ignoredLayer navmesh.LayerMask
}

func newActor(ctx *context.Context, player *context.Player, actorType battleAPI.ActorType,
	position *vec2.T, direction float64, params context.ActorParams) *actor {
	a := &actor{
		player:       player,
		actorType:    actorType,
		ctx:          ctx,
		event:        event.New(),
		motor:        newMotor(ctx, position, direction, params.AccelMaxSpeed(), params.AccelDuration()),
		stageAgent:   ctx.Stage.CreateAgent(21, position),
		ignoredLayer: player.TeamLayer,
	}

	switch actorType {
	case battleAPI.ActorType_Submarine:
		a.stageAgent.SetLayer(context.LayerSubmarine)
	case battleAPI.ActorType_Torpedo:
		a.stageAgent.SetLayer(context.LayerTorpedo)
	}
	a.stageAgent.SetLayer(a.Player().TeamLayer)
	a.stageAgent.SetCollideHandler(a.onStageAgentCollide)
	return a
}

func (a *actor) ID() int64                 { return a.stageAgent.ID() }
func (a *actor) Player() *context.Player   { return a.player }
func (a *actor) Type() battleAPI.ActorType { return a.actorType }
func (a *actor) Event() *event.Emitter     { return a.event }

func (a *actor) IsDestroyed() bool             { return a.isDestroyed }
func (a *actor) Movement() *battleAPI.Movement { return a.motor.toAPIType(a.ID()) }
func (a *actor) Position() *vec2.T             { return a.stageAgent.Position() }
func (a *actor) Direction() float64            { return a.motor.direction }
func (a *actor) IsAccelerating() bool          { return a.motor.accelerator.isAccelerating }

func (a *actor) Destroy() {
	a.isDestroyed = true
	a.stageAgent.Destroy()
	a.ctx.Event.Emit(event.ActorDestroy, a)
}

func (a *actor) BeforeUpdate() {
	position := a.motor.position()

	if a.player.AI == nil {
		a.stageAgent.Move(position, a.ignoredLayer)
	} else {
		a.stageAgent.Warp(position)
	}
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}

func (a *actor) accelerate(direction float64) {
	a.motor.accelerate()
	a.motor.turn(direction)
	a.ctx.Event.Emit(event.ActorMove, a)
}

func (a *actor) brake(direction float64) {
	a.motor.brake()
	a.motor.turn(direction)
	a.ctx.Event.Emit(event.ActorMove, a)
}

func (a *actor) turn(direction float64) {
	a.motor.turn(direction)
	a.ctx.Event.Emit(event.ActorMove, a)
}

func (a *actor) idle() {
	a.motor.idle(a.stageAgent.Position())
	a.ctx.Event.Emit(event.ActorMove, a)
}

func (a *actor) onStageAgentCollide(obj navmesh.Object, point vec2.T) {
	if a.IsDestroyed() {
		return
	}
	if obj == nil {
		a.event.Emit(event.ActorCollideWithStage, point)
	} else if other := a.ctx.Actor(obj.ID()); other != nil {
		a.event.Emit(event.ActorCollideWithOtherActor, other, point)
	}
}
