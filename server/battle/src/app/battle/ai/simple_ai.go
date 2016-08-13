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
func NewSimpleAI(ctx *context.Context) *SimpleAI {
	return &SimpleAI{ai: newAI(ctx)}
}

// Update the SimpleAI.
func (a *SimpleAI) Update(submarine context.Actor) {
	if !a.navigator.isStarted() {
		nextDest := a.nextDest(submarine.User().StartPosition)
		path := a.ctx.Stage.FindPath(submarine.Position(), nextDest)
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

func (a *SimpleAI) nextDest(startPosition *vec2.T) *vec2.T {
	var dest *vec2.T
	if a.isNextDestStartPosition {
		dest = &vec2.T{startPosition[0], startPosition[1]}
	} else {
		dest = &vec2.T{-startPosition[0], -startPosition[1]}
	}
	a.isNextDestStartPosition = !a.isNextDestStartPosition
	return dest
}
