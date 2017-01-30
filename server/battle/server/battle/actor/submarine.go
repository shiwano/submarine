package actor

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/actor/component"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/logger"

	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
	isUsingPinger bool
	timer         *component.Timer
	equipment     *component.Equipment
}

// NewSubmarine creates a submarine.
func NewSubmarine(ctx context.Context, user *context.Player) context.Actor {
	s := &submarine{
		actor: newActor(ctx, user, user.SubmarineParams, user.StartPosition, 0),
	}
	s.timer = component.NewTimer(ctx.Now())
	s.equipment = component.NewEquipment(s.ID(), user.SubmarineParams)

	s.event.AddCollideWithOtherActorEventListener(s.onCollideWithOtherActor)
	s.event.AddCollideWithStageEventListener(s.onCollideWithStage)
	s.event.AddAccelerationRequestEventListener(s.onAccelerationRequest)
	s.event.AddBrakeRequestEventListener(s.onBrakeRequest)
	s.event.AddTurnRequestEventListener(s.onTurnRequest)
	s.event.AddTorpedoRequestEventListener(s.onTorpedoRequest)
	s.event.AddPingerRequestEventListener(s.onPingerRequest)
	s.event.AddWatcherRequestEventListener(s.onWatcherRequest)
	s.event.AddUserLeaveEventListener(s.onUserLeave)

	s.ctx.Event().AddActorUsePingerEventListener(s.onActorUsePinger)
	s.ctx.Event().EmitActorCreateEvent(s)
	return s
}

func (s *submarine) Update() {
	s.timer.Update(s.ctx.Now())

	if s.player.AI != nil {
		s.player.AI.Update(s)
	}
}

func (s *submarine) OnDestroy() {
	if s.isUsingPinger {
		s.finishToUsePinger()
	}
}

func (s *submarine) Submarine(layer navmesh.LayerMask) *battleAPI.ActorSubmarineObject {
	e := &battleAPI.ActorSubmarineObject{
		IsUsingPinger: s.isUsingPinger,
	}
	if layer == s.player.TeamLayer {
		e.Equipment = s.equipment.ToAPIType()
	}
	return e
}

func (s *submarine) onCollideWithOtherActor(actor context.Actor, point vec2.T) {
	if actor.Player() != s.player {
		s.idle()
	}
}

func (s *submarine) onCollideWithStage(point vec2.T) {
	s.idle()
}

func (s *submarine) onAccelerationRequest(m *battleAPI.AccelerationRequestObject) {
	s.accelerate(m.Direction)
}

func (s *submarine) onBrakeRequest(m *battleAPI.BrakeRequestObject) {
	s.brake(m.Direction)
}

func (s *submarine) onTurnRequest(m *battleAPI.TurnRequestObject) {
	s.turn(m.Direction)
}

func (s *submarine) onTorpedoRequest(m *battleAPI.TorpedoRequestObject) {
	if s.equipment.TryConsumeTorpedo(s.ctx.Now()) {
		logger.Log.Debugf("%v shoots a torpedo", s)
		s.ctx.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		normalizedVelocity := s.motor.NormalizedVelocity()
		startOffsetValue := s.stageAgent.SizeRadius() * s.player.TorpedoParams.StartOffsetDistance
		startPoint := normalizedVelocity.Scale(startOffsetValue).Add(s.Position())
		NewTorpedo(s.ctx, s.player, startPoint, s.motor.Direction())
	}
}

func (s *submarine) onPingerRequest(m *battleAPI.PingerRequestObject) {
	if !s.isUsingPinger && s.equipment.Pinger.TryConsume(s.ctx.Now()) {
		logger.Log.Debugf("%v uses pinger", s)
		s.ctx.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		s.isUsingPinger = true
		s.ctx.Event().EmitActorUsePingerEvent(s, false)
		s.timer.Register(s.player.SubmarineParams.PingerIntervalSeconds, s.finishToUsePinger)
	}
}

func (s *submarine) onWatcherRequest(m *battleAPI.WatcherRequestObject) {
	if s.equipment.Watcher.TryConsume(s.ctx.Now()) {
		logger.Log.Debugf("%v uses watcher", s)
		s.ctx.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		NewWatcher(s.ctx, s.player, s.Position(), s.motor.Direction())
	}
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.Direction())
}

func (s *submarine) onActorUsePinger(a context.Actor, finished bool) {
	pingerTeam := a.Player().TeamLayer
	if s.Player().TeamLayer == pingerTeam {
		return
	}
	s.visibility.Set(pingerTeam, !finished)
}

func (s *submarine) finishToUsePinger() {
	if s.isUsingPinger {
		s.ctx.Event().EmitActorUsePingerEvent(s, true)
		s.isUsingPinger = false
	}
}
