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
	message *battleAPI.AccelerationRequestObject
}

type brakeRequestEvent struct {
	message *battleAPI.BrakeRequestObject
}

type turnRequestEvent struct {
	message *battleAPI.TurnRequestObject
}

type torpedoRequestEvent struct {
	message *battleAPI.TorpedoRequestObject
}

type pingerRequestEvent struct {
	message *battleAPI.PingerRequestObject
}

type watcherRequestEvent struct {
	message *battleAPI.WatcherRequestObject
}

type userLeaveEvent struct {
}
