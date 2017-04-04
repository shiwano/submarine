package ai

import (
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
)

type ai struct {
	scn       scene.Scene
	navigator *navigator
}

func newAI(scn scene.Scene) *ai {
	return &ai{
		scn:       scn,
		navigator: new(navigator),
	}
}

// Overridable methods.
func (a *ai) Update(submarine scene.Actor) {}

func (a *ai) accelerateActor(actor scene.Actor, dir float64) {
	actor.Event().EmitAccelerationRequestEvent(&battleAPI.AccelerationRequest{Direction: dir})
}

func (a *ai) brakeActor(actor scene.Actor, dir float64) {
	actor.Event().EmitBrakeRequestEvent(&battleAPI.BrakeRequest{Direction: dir})
}
