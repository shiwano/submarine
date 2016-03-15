package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

// Container holds the actors.
type Container struct {
	actors             []Actor
	actorsByID         map[int64]Actor
	submarinesByUserID map[int64]Actor
	context            *Context
}

func newContainer(battleContext *Context) *Container {
	c := &Container{
		actors:             make([]Actor, 0),
		actorsByID:         make(map[int64]Actor),
		submarinesByUserID: make(map[int64]Actor),
		context:            battleContext,
	}
	c.context.Event.On(event.ActorCreate, c.onActorCreate)
	c.context.Event.On(event.ActorDestroy, c.onActorDestroy)
	return c
}

func (c *Container) onActorCreate(actor Actor) {
	c.actorsByID[actor.ID()] = actor
	c.actors = append(c.actors, actor)
	if actor.ActorType() == battle.ActorType_Submarine {
		c.submarinesByUserID[actor.UserID()] = actor
	}
	actor.Start()
	c.context.Event.Emit(event.ActorAdd, actor)
}

func (c *Container) onActorDestroy(rawActor Actor) {
	actor := c.actorsByID[rawActor.ID()]
	if actor == nil {
		return
	}

	delete(c.actorsByID, actor.ID())
	c.actors = make([]Actor, len(c.actorsByID))
	for _, a := range c.actors {
		if a != actor {
			c.actors = append(c.actors, a)
		}
	}
	if actor.ActorType() == battle.ActorType_Submarine {
		delete(c.submarinesByUserID, actor.UserID())
	}
	actor.OnDestroy()
	c.context.Event.Emit(event.ActorRemove, actor)
}
