package battle

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi/core"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
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

func (g *Gateway) outputStart(receivers context.PlayerSlice, startedAt time.Time) {
	g.Output <- &GatewayOutput{
		UserIDs: receivers.SelectInt64(func(p *context.Player) int64 { return p.ID }),
		Message: &battleAPI.Start{
			StartedAt: currentmillis.Millis(startedAt),
		},
	}
}

func (g *Gateway) outputFinish(winnerUserID *int64, finishedAt time.Time) {
	g.Output <- &GatewayOutput{
		Message: &battleAPI.Finish{
			WinnerUserId: winnerUserID,
			FinishedAt:   currentmillis.Millis(finishedAt),
		},
		IsFinishMessage: true,
	}
}

func (g *Gateway) outputActor(receiversByTeam context.PlayersByTeam, actor context.Actor) {
	for teamLayer, receivers := range receiversByTeam {
		g.Output <- &GatewayOutput{
			UserIDs: receivers.SelectInt64(func(p *context.Player) int64 { return p.ID }),
			Message: &battleAPI.Actor{
				Id:        actor.ID(),
				UserId:    actor.Player().ID,
				Type:      actor.Type(),
				Movement:  actor.Movement(),
				IsVisible: actor.IsVisibleFrom(teamLayer),
			},
		}
	}
}

func (g *Gateway) outputVisibility(receiversByTeam context.PlayersByTeam, actor context.Actor,
	targetTeamLayer navmesh.LayerMask) {
	for teamLayer, receivers := range receiversByTeam {
		if teamLayer == targetTeamLayer {
			g.Output <- &GatewayOutput{
				UserIDs: receivers.SelectInt64(func(p *context.Player) int64 { return p.ID }),
				Message: &battleAPI.Visibility{
					ActorId:   actor.ID(),
					IsVisible: actor.IsVisibleFrom(teamLayer),
					Movement:  actor.Movement(),
				},
			}
		}
	}
}

func (g *Gateway) outputMovement(actor context.Actor) {
	g.Output <- &GatewayOutput{
		Message: actor.Movement(),
	}
}

func (g *Gateway) outputDestruction(actor context.Actor) {
	g.Output <- &GatewayOutput{
		Message: &battleAPI.Destruction{ActorId: actor.ID()},
	}
}

type gatewayInput struct {
	userID  int64
	message typhenapi.Type
}

// GatewayOutput represents a battle output.
type GatewayOutput struct {
	UserIDs         []int64
	Message         typhenapi.Type
	IsFinishMessage bool
}
