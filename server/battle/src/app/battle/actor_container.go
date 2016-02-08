package battle

import (
	"app/typhenapi/type/submarine/battle"
)

// ActorContainer creates actors, and holds the created.
type ActorContainer struct {
	nextActorID int64
	actors      map[int64]Actor
	submarines  map[int64]*Submarine
}

func newActorContainer() *ActorContainer {
	return &ActorContainer{
		nextActorID: 1,
		actors:      make(map[int64]Actor),
		submarines:  make(map[int64]*Submarine),
	}
}

func (c *ActorContainer) buildActor(userID int64, actorType battle.ActorType) *actor {
	nextActorID := c.nextActorID
	c.nextActorID++
	return &actor{
		id:        nextActorID,
		userID:    userID,
		actorType: actorType,
	}
}

func (c *ActorContainer) createSubmarine(userID int64) *Submarine {
	submarine := &Submarine{
		actor: c.buildActor(userID, battle.ActorType_Submarine),
	}
	c.actors[submarine.ID()] = submarine
	c.submarines[submarine.ID()] = submarine
	return submarine
}

func (c *ActorContainer) getActor(actorID int64) Actor {
	return c.actors[actorID]
}

func (c *ActorContainer) removeActor(actorID int64) {
	delete(c.actors, actorID)
}
