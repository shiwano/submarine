package ai

import (
	"app/battle/context"
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"
)

// AI represents a battle artificial intelligence.
type AI interface {
	Update(submarine context.Actor)
}

type ai struct {
	ctx       *context.Context
	navigator *navigator
}

func newAI(ctx *context.Context) *ai {
	return &ai{
		ctx:       ctx,
		navigator: new(navigator),
	}
}

// Overridable methods.
func (a *ai) Update(submarine context.Actor) {}

func (a *ai) accelerateActor(actor context.Actor, dir float64) {
	a.ctx.Event.Emit(event.AccelerationRequest, &battleAPI.AccelerationRequestObject{Direction: dir})
}

func (a *ai) brakeActor(actor context.Actor, dir float64) {
	a.ctx.Event.Emit(event.BrakeRequest, &battleAPI.BrakeRequestObject{Direction: dir})
}
