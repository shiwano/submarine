package battle

import (
	"app/typhenapi/type/submarine/battle"
)

// Actor represents an actor in the battle.
type Actor struct {
	id        int
	userID    int64
	actorType battle.ActorType
	position  *battle.Vector
}
