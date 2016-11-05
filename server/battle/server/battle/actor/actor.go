package actor

import (
	"fmt"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/logger"

	"github.com/ungerik/go3d/float64/vec2"
)

type actor struct {
	player             *context.Player
	actorType          battleAPI.ActorType
	ctx                *context.Context
	event              *context.ActorEventEmitter
	isDestroyed        bool
	motor              *motor
	stageAgent         *navmesh.Agent
	ignoredLayer       navmesh.LayerMask
	hasLight           bool
	isAlwaysVisible    bool
	visibilitiesByTeam map[navmesh.LayerMask]bool
}

func newActor(ctx *context.Context, player *context.Player, params context.ActorParams,
	position *vec2.T, direction float64) *actor {
	a := &actor{
		player:             player,
		actorType:          params.Type(),
		ctx:                ctx,
		event:              context.NewActorEventEmitter(),
		motor:              newMotor(ctx, position, direction, params.AccelMaxSpeed(), params.AccelDuration()),
		stageAgent:         ctx.Stage.CreateAgent(21, position),
		ignoredLayer:       player.TeamLayer,
		hasLight:           params.HasLight(),
		isAlwaysVisible:    params.IsAlwaysVisible(),
		visibilitiesByTeam: make(map[navmesh.LayerMask]bool),
	}

	switch params.Type() {
	case battleAPI.ActorType_Submarine:
		a.stageAgent.SetLayer(context.LayerSubmarine)
	case battleAPI.ActorType_Torpedo:
		a.stageAgent.SetLayer(context.LayerTorpedo)
	}
	a.stageAgent.SetLayer(a.Player().TeamLayer)
	return a
}

func (a *actor) String() string {
	return fmt.Sprintf("%v's %v(%v)", a.player, a.actorType, a.stageAgent.ID())
}

func (a *actor) ID() int64                         { return a.stageAgent.ID() }
func (a *actor) Player() *context.Player           { return a.player }
func (a *actor) Type() battleAPI.ActorType         { return a.actorType }
func (a *actor) Event() *context.ActorEventEmitter { return a.event }

func (a *actor) IsDestroyed() bool             { return a.isDestroyed }
func (a *actor) Movement() *battleAPI.Movement { return a.motor.toAPIType(a.ID()) }
func (a *actor) Position() *vec2.T             { return a.stageAgent.Position() }
func (a *actor) Direction() float64            { return a.motor.direction }
func (a *actor) IsAccelerating() bool          { return a.motor.accelerator.isAccelerating }

func (a *actor) IsVisibleFrom(layer navmesh.LayerMask) bool {
	if a.isAlwaysVisible || a.player.TeamLayer == layer {
		return true
	} else if isVisible, ok := a.visibilitiesByTeam[layer]; ok {
		return isVisible
	}
	return false
}

func (a *actor) Destroy() {
	a.isDestroyed = true
	a.stageAgent.Destroy()
	a.ctx.Event.EmitActorDestroyEvent(a)
}

func (a *actor) BeforeUpdate() {
	position := a.motor.position()

	if hitInfo := a.stageAgent.Move(position, a.ignoredLayer); hitInfo != nil {
		a.onStageAgentCollide(hitInfo.Object, hitInfo.Point)
	}

	if a.hasLight {
		a.ctx.SightsByTeam[a.Player().TeamLayer].PutLight(a.Position())
	}
}

func (a *actor) AfterUpdate() {
	if !a.isAlwaysVisible {
		for _, teamLayer := range context.TeamLayers {
			newVisibility := a.ctx.SightsByTeam[teamLayer].IsLitPoint(a.Position())
			oldVisibility := a.visibilitiesByTeam[teamLayer]
			a.visibilitiesByTeam[teamLayer] = newVisibility

			if newVisibility != oldVisibility {
				a.ctx.Event.EmitActorChangeVisibilityEvent(a, teamLayer)
			}
		}
	}
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}

func (a *actor) accelerate(direction float64) {
	logger.Log.Debugf("%v accelerates to %v", a, direction)
	a.motor.accelerate(a.stageAgent.Position())
	a.motor.turn(a.stageAgent.Position(), direction)
	a.ctx.Event.EmitActorMoveEvent(a)
}

func (a *actor) brake(direction float64) {
	logger.Log.Debugf("%v brakes", a)
	a.motor.brake(a.stageAgent.Position())
	a.motor.turn(a.stageAgent.Position(), direction)
	a.ctx.Event.EmitActorMoveEvent(a)
}

func (a *actor) turn(direction float64) {
	logger.Log.Debugf("%v turns to %v", a, direction)
	a.motor.turn(a.stageAgent.Position(), direction)
	a.ctx.Event.EmitActorMoveEvent(a)
}

func (a *actor) idle() {
	a.motor.idle(a.stageAgent.Position())
	a.ctx.Event.EmitActorMoveEvent(a)
}

func (a *actor) onStageAgentCollide(obj navmesh.Object, point vec2.T) {
	if a.IsDestroyed() {
		return
	}
	if obj == nil {
		logger.Log.Debugf("%v collided with stage", a)
		a.event.EmitCollideWithStageEvent(point)
	} else if other := a.ctx.Actor(obj.ID()); other != nil {
		logger.Log.Debugf("%v collided with %v", a, other)
		a.event.EmitCollideWithOtherActorEvent(other, point)
	}
}
