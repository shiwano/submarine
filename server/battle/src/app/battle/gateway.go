package battle

import (
	"app/battle/context"
	"app/currentmillis"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	"time"
)

// Gateway represents a battle input/output.
type Gateway struct {
	Output chan typhenapi.Type
	input  chan *gatewayInput
}

func newGateway() *Gateway {
	return &Gateway{
		Output: make(chan typhenapi.Type, 256),
		input:  make(chan *gatewayInput, 256),
	}
}

// InputMessage sends the user's message to the input channel.
func (g *Gateway) InputMessage(userID int64, message typhenapi.Type) {
	g.input <- &gatewayInput{
		userID:  userID,
		message: message,
	}
}

func (g *Gateway) outputStart(startedAt time.Time) {
	g.Output <- &battle.Start{
		StartedAt: currentmillis.Milliseconds(startedAt),
	}
}

func (g *Gateway) outputFinish(winnerUserID int64, finishedAt time.Time) {
	g.Output <- &battle.Finish{
		WinnerUserId: winnerUserID,
		FinishedAt:   currentmillis.Milliseconds(finishedAt),
	}
}

func (g *Gateway) outputActor(actor context.Actor) {
	g.Output <- &battle.Actor{
		Id:       actor.ID(),
		UserId:   actor.UserID(),
		Type:     actor.ActorType(),
		Movement: actor.Movement(),
	}
}

func (g *Gateway) outputMovement(actor context.Actor) {
	g.Output <- actor.Movement()
}

type gatewayInput struct {
	userID  int64
	message typhenapi.Type
}
