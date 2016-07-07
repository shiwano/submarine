package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"github.com/ungerik/go3d/float64/vec2"
)

type torpedo struct {
	*actor
}

// NewTorpedo creates a torpedo.
func NewTorpedo(battleContext *context.Context, user *context.User, position *vec2.T, direction float64) context.Actor {
	t := &torpedo{
		actor: newActor(battleContext, user, battle.ActorType_Torpedo, position, direction, user.TorpedoParams),
	}
	t.event.On(event.ActorCollideWithStage, t.onCollideWithStage)
	t.event.On(event.ActorCollideWithOtherActor, t.onCollideWithOtherActor)
	t.context.Event.Emit(event.ActorCreate, t)
	t.accelerate(direction)
	return t
}

func (t *torpedo) onCollideWithStage(point vec2.T) {
	t.Destroy()
}

func (t *torpedo) onCollideWithOtherActor(actor context.Actor, point vec2.T) {
	if actor.User() != t.User() {
		t.Destroy()

		if actor.Type() == battle.ActorType_Submarine {
			actor.Destroy()
		}
	}
}
