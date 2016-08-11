package ai

import (
	"app/battle/context"
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"
)

// AI represents a battle artificial intelligence.
type AI interface {
	Update()
}

type ai struct {
	ctx       context.Context
	user      *context.User
	navigator *navigator
}

func newAI(ctx context.Context, user *context.User) *ai {
	return &ai{
		ctx:       ctx,
		user:      user,
		navigator: new(navigator),
	}
}

// Overridable methods.
func (a *ai) Update() {}

func (a *ai) accelerateActor(actor context.Actor, dir float64) {
	a.ctx.Event.Emit(event.AccelerationRequest, &battleAPI.AccelerationRequestObject{Direction: dir})
}

func (a *ai) brakeActor(actor context.Actor, dir float64) {
	a.ctx.Event.Emit(event.BrakeRequest, &battleAPI.BrakeRequestObject{Direction: dir})
}
