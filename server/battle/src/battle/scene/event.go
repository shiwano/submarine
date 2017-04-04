//go:generate genem -n EventEmitter -o event_emitter.go $GOFILE

package scene

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

type actorCreateEvent struct {
	actor Actor
}

type actorDestroyEvent struct {
	actor Actor
}

type actorAddEvent struct {
	actor Actor
}

type actorMoveEvent struct {
	actor Actor
}

type actorRemoveEvent struct {
	actor Actor
}

type actorChangeVisibilityEvent struct {
	actor     Actor
	teamLayer navmesh.LayerMask
}

type actorUsePingerEvent struct {
	actor    Actor
	finished bool
}

type actorUpdateEquipmentEvent struct {
	actor     Actor
	equipment *battleAPI.Equipment
}
