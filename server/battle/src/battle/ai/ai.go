package ai

import (
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/context"
)

type ai struct {
	ctx       context.Context
	navigator *navigator
}

func newAI(ctx context.Context) *ai {
	return &ai{
		ctx:       ctx,
		navigator: new(navigator),
	}
}

// Overridable methods.
func (a *ai) Update(submarine context.Actor) {}

func (a *ai) accelerateActor(actor context.Actor, dir float64) {
	actor.Event().EmitAccelerationRequestEvent(&battleAPI.AccelerationRequest{Direction: dir})
}

func (a *ai) brakeActor(actor context.Actor, dir float64) {
	actor.Event().EmitBrakeRequestEvent(&battleAPI.BrakeRequest{Direction: dir})
}
