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
func NewSubmarine(battleContext *context.Context, userID int64) context.Actor {
	s := &submarine{
		actor: newActor(battleContext, userID, battle.ActorType_Submarine),
	}
	s.context.Event.EmitSync(event.ActorCreated, s)
	return s
}
