package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/logger"
	battleAPI "app/typhenapi/type/submarine/battle"
	"fmt"
	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(ctx *context.Context, user *context.Player) context.Actor {
	s := &submarine{
		actor: newActor(ctx, user, battleAPI.ActorType_Submarine, user.StartPosition, 0, user.SubmarineParams),
	}
	s.event.On(event.ActorCollideWithOtherActor, s.onCollideWithOtherActor)
	s.event.On(event.ActorCollideWithStage, s.onCollideWithStage)
	s.event.On(event.AccelerationRequest, s.onAccelerationRequest)
	s.event.On(event.BrakeRequest, s.onBrakeRequest)
	s.event.On(event.TurnRequest, s.onTurnRequest)
	s.event.On(event.TorpedoRequest, s.onTorpedoRequest)
	s.event.On(event.PingerRequest, s.onPingerRequest)
	s.event.On(event.UserLeave, s.onUserLeave)
	s.ctx.Event.Emit(event.ActorCreate, s)
	return s
}

func (s *submarine) String() string {
	return fmt.Sprintf("User(%v)'s submarine(%v)", s.Player().ID, s.ID())
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
	if s.player.AI == nil {
		s.idle()
	}
}

func (s *submarine) onAccelerationRequest(m *battleAPI.AccelerationRequestObject) {
	logger.Log.Debugf("%v accelerates to %v", s, m.Direction)
	s.accelerate(m.Direction)
}

func (s *submarine) onBrakeRequest(m *battleAPI.BrakeRequestObject) {
	logger.Log.Debugf("%v brakes", s)
	s.brake(m.Direction)
}

func (s *submarine) onTurnRequest(m *battleAPI.TurnRequestObject) {
	logger.Log.Debugf("%v turns to %v", s, m.Direction)
	s.turn(m.Direction)
}

func (s *submarine) onTorpedoRequest(m *battleAPI.TorpedoRequestObject) {
	logger.Log.Debugf("%v shoots a torpedo", s)
	s.shootTorpedo()
}

func (s *submarine) onPingerRequest(m *battleAPI.PingerRequestObject) {
	logger.Log.Debugf("%v uses pinger", s)
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.direction)
}

func (s *submarine) shootTorpedo() {
	p := s.motor.normalizedVelocity.Scaled(s.stageAgent.SizeRadius() * s.player.TorpedoParams.StartOffsetDistance)
	p.Add(s.Position())
	NewTorpedo(s.ctx, s.player, &p, s.motor.direction)
}
