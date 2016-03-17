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
	Output chan *GatewayOutput
	input  chan *gatewayInput
}

func newGateway() *Gateway {
	return &Gateway{
		Output: make(chan *GatewayOutput, 256),
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

func (g *Gateway) outputMessage(userIDs []int64, message typhenapi.Type) {
	g.Output <- &GatewayOutput{
		UserIDs: userIDs,
		Message: message,
	}
}

func (g *Gateway) outputStart(startedAt time.Time) {
	g.outputMessage(nil, &battle.Start{
		StartedAt: currentmillis.Milliseconds(startedAt),
	})
}

func (g *Gateway) outputFinish(winnerUserID int64, finishedAt time.Time) {
	g.outputMessage(nil, &battle.Finish{
		WinnerUserId: winnerUserID,
		FinishedAt:   currentmillis.Milliseconds(finishedAt),
	})
}

func (g *Gateway) outputActor(userIDs []int64, actor context.Actor) {
	g.outputMessage(userIDs, &battle.Actor{
		Id:       actor.ID(),
		UserId:   actor.UserID(),
		Type:     actor.Type(),
		Movement: actor.Movement(),
	})
}

func (g *Gateway) outputMovement(actor context.Actor) {
	g.outputMessage(nil, actor.Movement())
}

type gatewayInput struct {
	userID  int64
	message typhenapi.Type
}

// GatewayOutput represents a battle output.
type GatewayOutput struct {
	UserIDs []int64
	Message typhenapi.Type
}
