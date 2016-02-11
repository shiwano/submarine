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
}

func newGateway() *Gateway {
	return &Gateway{
		Output: make(chan interface{}, 256),
		Input:  make(chan interface{}, 256),
	}
}

func (g *Gateway) start(startedAt time.Time) {
	g.Output <- &battle.Start{
		StartedAt: currentmillis.ToMilliseconds(startedAt),
	}
}

func (g *Gateway) finish(winnerUserID int64, finishedAt time.Time) {
	g.Output <- &battle.Finish{
		WinnerUserId: winnerUserID,
		FinishedAt:   currentmillis.ToMilliseconds(finishedAt),
	}
}

func (g *Gateway) actor(actor Actor) {
	position := actor.Position()
	g.Output <- &battle.Actor{
		Id:       actor.ID(),
		UserId:   actor.UserID(),
		Type:     actor.ActorType(),
		Position: &position,
	}
}
