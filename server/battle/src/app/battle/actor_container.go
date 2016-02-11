package battle

import (
	"app/logger"
	"app/typhenapi/type/submarine/battle"
)

// ActorContainer creates actors, and holds the created.
type ActorContainer struct {
	nextActorID int64
	actors      map[int64]Actor
	submarines  map[int64]*Submarine
	context     *Context
}

func newActorContainer(context *Context) *ActorContainer {
	return &ActorContainer{
		nextActorID: 1,
		actors:      make(map[int64]Actor),
		submarines:  make(map[int64]*Submarine),
		context:     context,
	}
}

func (c *ActorContainer) buildActor(userID int64, actorType battle.ActorType) *actor {
	nextActorID := c.nextActorID
	c.nextActorID++
	return &actor{
		id:        nextActorID,
		userID:    userID,
		actorType: actorType,
		context:   c.context,
	}
}

func (c *ActorContainer) createSubmarine(userID int64) *Submarine {
	if _, ok := c.submarines[userID]; ok {
		logger.Log.Errorf("User(%v)'s submarine already exists", userID)
		return nil
	}
	submarine := &Submarine{
		actor: c.buildActor(userID, battle.ActorType_Submarine),
	}
	c.actors[submarine.ID()] = submarine
	c.submarines[submarine.UserID()] = submarine
	return submarine
}

func (c *ActorContainer) getSubmarineByUserID(userID int64) *Submarine {
	return c.submarines[userID]
}

func (c *ActorContainer) getActor(actorID int64) Actor {
	return c.actors[actorID]
}

func (c *ActorContainer) removeActor(actorID int64) {
	delete(c.actors, actorID)
}
