package battle

import (
	"app/battle/context"
	"app/typhenapi/core"
	"app/typhenapi/type/submarine/battle"
	"lib/currentmillis"
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

func (g *Gateway) outputStart(userIDs []int64, startedAt time.Time) {
	g.outputMessage(userIDs, &battle.Start{
		StartedAt: currentmillis.Millis(startedAt),
	})
}

func (g *Gateway) outputFinish(winnerUserID int64, finishedAt time.Time) {
	g.outputMessage(nil, &battle.Finish{
		WinnerUserId: winnerUserID,
		FinishedAt:   currentmillis.Millis(finishedAt),
	})
}

func (g *Gateway) outputActor(userIDs []int64, actor context.Actor) {
	g.outputMessage(userIDs, &battle.Actor{
		Id:       actor.ID(),
		UserId:   actor.User().ID,
		Type:     actor.Type(),
		Movement: actor.Movement(),
	})
}

func (g *Gateway) outputMovement(actor context.Actor) {
	g.outputMessage(nil, actor.Movement())
}

func (g *Gateway) outputDestruction(actor context.Actor) {
	g.outputMessage(nil, &battle.Destruction{ActorId: actor.ID()})
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
