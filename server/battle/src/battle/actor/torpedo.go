package actor

import (
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"

	"github.com/ungerik/go3d/float64/vec2"
)

type torpedo struct {
	*actor
}

func newTorpedo(scn scene.Scene, user *scene.Player, position *vec2.T, direction float64) scene.Actor {
	t := &torpedo{
		actor: newActor(scn, user, user.TorpedoParams, position, direction),
	}
	t.event.AddCollideWithStageEventListener(t.onCollideWithStage)
	t.event.AddCollideWithOtherActorEventListener(t.onCollideWithOtherActor)
	t.scene.Event().EmitActorCreateEvent(t)
	t.accelerate(direction)
	return t
}

func (t *torpedo) onCollideWithStage(point vec2.T) {
	t.Destroy()
}

func (t *torpedo) onCollideWithOtherActor(actor scene.Actor, point vec2.T) {
	t.Destroy()

	if actor.Type() == battleAPI.ActorType_Submarine {
		actor.Destroy()
	}
}
