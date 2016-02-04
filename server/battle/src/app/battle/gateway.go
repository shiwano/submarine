package battle

import (
	"app/typhenapi/type/submarine/battle"
)

// Gateway represents a battle input/output.
type Gateway struct {
	Output chan interface{}
	Input  chan interface{}
	Close  chan struct{}
}

func newGateway() *Gateway {
	return &Gateway{
		Output: make(chan interface{}, 256),
		Input:  make(chan interface{}, 256),
		Close:  make(chan struct{}, 1),
	}
}

func (g *Gateway) start(startedAt int64) {
	g.Output <- &battle.Start{
		StartedAt: startedAt,
	}
}

func (g *Gateway) finish(hasWon bool, finishedAt int64) {
	g.Output <- &battle.Finish{
		HasWon:     hasWon,
		FinishedAt: finishedAt,
	}
}
