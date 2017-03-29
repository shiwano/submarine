//go:generate genem -n ActorEventEmitter -o actor_event_emitter.go $GOFILE

package context

import (
	"github.com/ungerik/go3d/float64/vec2"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

type collideWithStageEvent struct {
	point vec2.T
}

type collideWithOtherActorEvent struct {
	actor Actor
	point vec2.T
}

type accelerationRequestEvent struct {
	message *battleAPI.AccelerationRequest
}

type brakeRequestEvent struct {
	message *battleAPI.BrakeRequest
}

type turnRequestEvent struct {
	message *battleAPI.TurnRequest
}

type torpedoRequestEvent struct {
	message *battleAPI.TorpedoRequest
}

type pingerRequestEvent struct {
	message *battleAPI.PingerRequest
}

type watcherRequestEvent struct {
	message *battleAPI.WatcherRequest
}

type userLeaveEvent struct {
}
