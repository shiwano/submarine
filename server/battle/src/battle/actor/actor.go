package actor

import (
	"fmt"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/actor/component"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
	"github.com/shiwano/submarine/server/battle/src/logger"

	"github.com/ungerik/go3d/float64/vec2"
)

type actor struct {
	player       *scene.Player
	actorType    battleAPI.ActorType
	scene        scene.Scene
	event        *scene.ActorEventEmitter
	isDestroyed  bool
	motor        *component.Motor
	stageAgent   *navmesh.Agent
	ignoredLayer navmesh.LayerMask
	hasLight     bool
	visibility   *component.Visibility
	isLitByTeam  map[navmesh.LayerMask]bool
}

func newActor(scn scene.Scene, player *scene.Player, params scene.ActorParams,
	position *vec2.T, direction float64) *actor {
	a := &actor{
		player:       player,
		actorType:    params.Type(),
		scene:        scn,
		event:        scene.NewActorEventEmitter(),
		motor:        component.NewMotor(scn, position, direction, params.AccelMaxSpeed(), params.AccelDuration()),
		stageAgent:   scn.Stage().CreateAgent(21, position),
		ignoredLayer: player.TeamLayer,
		hasLight:     params.HasLight(),
		visibility:   component.NewVisibility(),
		isLitByTeam:  make(map[navmesh.LayerMask]bool),
	}
	for _, l := range scene.TeamLayers {
		if params.IsAlwaysVisible() || a.player.TeamLayer == l {
			a.visibility.Set(l, true)
		}
	}
	a.visibility.ChangeHandler = a.onVisibilityChange

	switch params.Type() {
	case battleAPI.ActorType_Submarine:
		a.stageAgent.SetLayer(a.Player().TeamLayer | scene.LayerSubmarine)
	case battleAPI.ActorType_Torpedo:
		a.stageAgent.SetLayer(a.Player().TeamLayer | scene.LayerTorpedo)
	default:
		a.stageAgent.SetLayer(a.Player().TeamLayer)
	}
	return a
}

func (a *actor) String() string {
	return fmt.Sprintf("%v's %v(%v)", a.player, a.actorType, a.stageAgent.ID())
}

func (a *actor) ID() int64                                  { return a.stageAgent.ID() }
func (a *actor) Player() *scene.Player                      { return a.player }
func (a *actor) Type() battleAPI.ActorType                  { return a.actorType }
func (a *actor) Event() *scene.ActorEventEmitter            { return a.event }
func (a *actor) IsDestroyed() bool                          { return a.isDestroyed }
func (a *actor) Movement() *battleAPI.Movement              { return a.motor.ToAPIType(a.ID()) }
func (a *actor) Position() *vec2.T                          { return a.stageAgent.Position() }
func (a *actor) Direction() float64                         { return a.motor.Direction() }
func (a *actor) IsAccelerating() bool                       { return a.motor.IsAccelerating() }
func (a *actor) IsVisibleFrom(layer navmesh.LayerMask) bool { return a.visibility.IsVisibleFrom(layer) }

func (a *actor) Submarine(navmesh.LayerMask) *battleAPI.ActorSubmarineObject { return nil }

func (a *actor) BeforeUpdate() {
	position := a.motor.Position()
	if hitInfo := a.stageAgent.Move(position, a.ignoredLayer); hitInfo != nil {
		a.onStageAgentCollide(hitInfo.Object, hitInfo.Point)
	}
	if a.hasLight {
		a.scene.SightsByTeam()[a.Player().TeamLayer].PutLight(a.Position())
	}
}

func (a *actor) AfterUpdate() {
	a.refreshVisibilities()
}

// Overridable methods.
func (a *actor) Start()     {}
func (a *actor) Update()    {}
func (a *actor) OnDestroy() {}

func (a *actor) Destroy() {
	a.isDestroyed = true
	a.stageAgent.Destroy()
	a.scene.Event().EmitActorDestroyEvent(a)
}

func (a *actor) accelerate(direction float64) {
	logger.Log.Debugf("%v accelerates to %v", a, direction)
	a.motor.Accelerate(a.stageAgent.Position())
	a.motor.Turn(a.stageAgent.Position(), direction)
	a.scene.Event().EmitActorMoveEvent(a)
}

func (a *actor) brake(direction float64) {
	logger.Log.Debugf("%v brakes", a)
	a.motor.Brake(a.stageAgent.Position())
	a.motor.Turn(a.stageAgent.Position(), direction)
	a.scene.Event().EmitActorMoveEvent(a)
}

func (a *actor) turn(direction float64) {
	logger.Log.Debugf("%v turns to %v", a, direction)
	a.motor.Turn(a.stageAgent.Position(), direction)
	a.scene.Event().EmitActorMoveEvent(a)
}

func (a *actor) idle() {
	a.motor.Idle(a.stageAgent.Position())
	a.scene.Event().EmitActorMoveEvent(a)
}

func (a *actor) onStageAgentCollide(obj navmesh.Object, point vec2.T) {
	if a.IsDestroyed() {
		return
	}
	if obj == nil {
		logger.Log.Debugf("%v collided with stage", a)
		a.event.EmitCollideWithStageEvent(point)
	} else if other, ok := a.scene.Actor(obj.ID()); ok {
		logger.Log.Debugf("%v collided with %v", a, other)
		a.event.EmitCollideWithOtherActorEvent(other, point)
	}
}

func (a *actor) onVisibilityChange(layer navmesh.LayerMask) {
	a.scene.Event().EmitActorChangeVisibilityEvent(a, layer)
}

func (a *actor) refreshVisibilities() {
	for _, l := range scene.TeamLayers {
		isLit := a.scene.SightsByTeam()[l].IsLitPoint(a.Position())
		if isLit != a.isLitByTeam[l] {
			a.isLitByTeam[l] = isLit
			a.visibility.Set(l, isLit)
		}
	}
}
