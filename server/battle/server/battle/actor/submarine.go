package actor

import (
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/logger"

	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(ctx *context.Context, user *context.Player) context.Actor {
	s := &submarine{
		actor: newActor(ctx, user, user.SubmarineParams, user.StartPosition, 0),
	}
	s.event.AddCollideWithOtherActorEventListener(s.onCollideWithOtherActor)
	s.event.AddCollideWithStageEventListener(s.onCollideWithStage)
	s.event.AddAccelerationRequestEventListener(s.onAccelerationRequest)
	s.event.AddBrakeRequestEventListener(s.onBrakeRequest)
	s.event.AddTurnRequestEventListener(s.onTurnRequest)
	s.event.AddTorpedoRequestEventListener(s.onTorpedoRequest)
	s.event.AddPingerRequestEventListener(s.onPingerRequest)
	s.event.AddUserLeaveEventListener(s.onUserLeave)
	s.ctx.Event.EmitActorCreateEvent(s)
	return s
}

func (s *submarine) Update() {
	if s.player.AI != nil {
		s.player.AI.Update(s)
	}
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
	s.shootTorpedo()
}

func (s *submarine) onPingerRequest(m *battleAPI.PingerRequestObject) {
	logger.Log.Debugf("%v uses pinger", s)
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.direction)
}

func (s *submarine) shootTorpedo() {
	logger.Log.Debugf("%v shoots a torpedo", s)
	p := s.motor.normalizedVelocity.Scaled(s.stageAgent.SizeRadius() * s.player.TorpedoParams.StartOffsetDistance)
	p.Add(s.Position())
	NewTorpedo(s.ctx, s.player, &p, s.motor.direction)
}
