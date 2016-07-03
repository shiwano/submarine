package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(battleContext *context.Context, user *context.User) context.Actor {
	s := &submarine{
		actor: newActor(battleContext, user, battle.ActorType_Submarine, user.StartPosition, 0,
			user.SubmarineAccelerationMaxSpeed, user.SubmarineAccelerationDuration),
	}
	s.event.On(event.ActorCollide, s.onCollide)
	s.event.On(event.AccelerationRequest, s.onAccelerationRequest)
	s.event.On(event.BrakeRequest, s.onBrakeRequest)
	s.event.On(event.TurnRequest, s.onTurnRequest)
	s.event.On(event.TorpedoRequest, s.onTorpedoRequest)
	s.event.On(event.UserLeave, s.onUserLeave)
	s.context.Event.Emit(event.ActorCreate, s)
	return s
}

func (s *submarine) onCollide(actor context.Actor, point vec2.T) {
	if actor == nil || actor.User() != s.User() {
		s.idle()
	}
}

func (s *submarine) onAccelerationRequest(message *battle.AccelerationRequestObject) {
	s.accelerate(message.Direction)
}

func (s *submarine) onBrakeRequest(message *battle.BrakeRequestObject) {
	s.brake(message.Direction)
}

func (s *submarine) onTurnRequest(message *battle.TurnRequestObject) {
	s.turn(message.Direction)
}

func (s *submarine) onTorpedoRequest(message *battle.TorpedoRequestObject) {
	s.shootTorpedo()
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.direction)
}

func (s *submarine) shootTorpedo() {
	p := s.motor.normalizedVelocity.Scaled(s.stageAgent.SizeRadius() * s.user.TorpedoStartOffsetLength)
	p.Add(s.Position())
	NewTorpedo(s.context, s.user, &p, s.motor.direction)
}
