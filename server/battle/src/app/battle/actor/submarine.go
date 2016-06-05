package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(battleContext *context.Context, user *context.User) context.Actor {
	s := &submarine{
		actor: newActor(battleContext, user, battle.ActorType_Submarine),
	}
	s.event.On(event.AccelerationRequest, s.onAccelerationRequest)
	s.event.On(event.BrakeRequest, s.onBrakeRequest)
	s.event.On(event.TurnRequest, s.onTurnRequest)
	s.event.On(event.UserLeave, s.onUserLeave)
	s.context.Event.Emit(event.ActorCreate, s)
	return s
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

func (s *submarine) onUserLeave() {
	s.brake(s.motor.direction)
}
