package actor

import (
	"app/battle/context"
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(ctx *context.Context, user *context.User) context.Actor {
	s := &submarine{
		actor: newActor(ctx, user, battleAPI.ActorType_Submarine, user.StartPosition, 0, user.SubmarineParams),
	}
	s.event.On(event.ActorCollideWithOtherActor, s.onCollideWithOtherActor)
	s.event.On(event.AccelerationRequest, s.onAccelerationRequest)
	s.event.On(event.BrakeRequest, s.onBrakeRequest)
	s.event.On(event.TurnRequest, s.onTurnRequest)
	s.event.On(event.TorpedoRequest, s.onTorpedoRequest)
	s.event.On(event.UserLeave, s.onUserLeave)
	s.ctx.Event.Emit(event.ActorCreate, s)
	return s
}

func (s *submarine) onCollideWithOtherActor(actor context.Actor, point vec2.T) {
	if actor.User() != s.User() {
		s.idle()
	}
}

func (s *submarine) onAccelerationRequest(message *battleAPI.AccelerationRequestObject) {
	s.accelerate(message.Direction)
}

func (s *submarine) onBrakeRequest(message *battleAPI.BrakeRequestObject) {
	s.brake(message.Direction)
}

func (s *submarine) onTurnRequest(message *battleAPI.TurnRequestObject) {
	s.turn(message.Direction)
}

func (s *submarine) onTorpedoRequest(message *battleAPI.TorpedoRequestObject) {
	s.shootTorpedo()
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.direction)
}

func (s *submarine) shootTorpedo() {
	p := s.motor.normalizedVelocity.Scaled(s.stageAgent.SizeRadius() * s.user.TorpedoParams.StartOffsetDistance)
	p.Add(s.Position())
	NewTorpedo(s.ctx, s.user, &p, s.motor.direction)
}
