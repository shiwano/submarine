package ai

import (
	"app/battle/context"
	"app/util"
	"github.com/ungerik/go3d/float64/vec2"
)

// SimpleAI represents a simple battle AI.
type SimpleAI struct {
	*ai
	isNextDestStartPosition bool
}

// NewSimpleAI creates a SimpleAI.
func NewSimpleAI(ctx context.Context, user *context.User) *SimpleAI {
	return &SimpleAI{ai: newAI(ctx, user)}
}

// Update the SimpleAI.
func (a *SimpleAI) Update() {
	submarine := a.ctx.SubmarineByUserID(a.user.ID)

	if !a.navigator.isStarted() {
		path := a.ctx.Stage.FindPath(submarine.Position(), a.nextDest())
		a.navigator.start(path, submarine.Position())
	}

	if ok, dir := a.navigator.navigate(submarine.Position()); ok {
		if !util.EqualFloats(dir, submarine.Direction()) {
			a.accelerateActor(submarine, dir)
		}
	} else {
		if submarine.IsAccelerating() {
			a.brakeActor(submarine, submarine.Direction())
		}
	}
}

func (a *SimpleAI) nextDest() *vec2.T {
	var dest *vec2.T
	if a.isNextDestStartPosition {
		dest = &vec2.T{a.user.StartPosition[0], a.user.StartPosition[1]}
	} else {
		dest = &vec2.T{-a.user.StartPosition[0], -a.user.StartPosition[1]}
	}
	a.isNextDestStartPosition = !a.isNextDestStartPosition
	return dest
}
