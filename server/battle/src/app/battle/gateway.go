package battle

import (
	"app/currentmillis"
	"app/typhenapi/type/submarine/battle"
	"time"
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

func (g *Gateway) start(startedAt time.Time) {
	g.Output <- &battle.Start{
		StartedAt: currentmillis.ToMilliseconds(startedAt),
	}
}

func (g *Gateway) finish(hasWon bool, finishedAt time.Time) {
	g.Output <- &battle.Finish{
		HasWon:     hasWon,
		FinishedAt: currentmillis.ToMilliseconds(finishedAt),
	}
}

func (g *Gateway) actor(actor Actor) {
	g.Output <- &battle.Actor{
		Id:       actor.ID(),
		UserId:   actor.UserID(),
		Type:     actor.ActorType(),
		Position: actor.Position(),
	}
}
