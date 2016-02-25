package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

// Container holds the actors.
type Container struct {
	actors     map[int64]Actor
	submarines map[int64]Actor
	context    *Context
}

func newContainer(battleContext *Context) *Container {
	c := &Container{
		actors:     make(map[int64]Actor),
		submarines: make(map[int64]Actor),
		context:    battleContext,
	}
	c.context.Event.On(event.ActorCreated, c.onActorCreate)
	c.context.Event.On(event.ActorDestroyed, c.onActorDestroy)
	return c
}

// UpdateActors updates all actors.
func (c *Container) UpdateActors() {
	for _, actor := range c.actors {
		actor.Update()
	}
}

// SubmarineByUserID returns the submarine which has the given actor id.
func (c *Container) SubmarineByUserID(userID int64) Actor {
	return c.submarines[userID]
}

// Actor returns the actor that has the actor id.
func (c *Container) Actor(actorID int64) Actor {
	return c.actors[actorID]
}

// HasActor determines whether the specified actor exists.
func (c *Container) HasActor(actorID int64) bool {
	_, ok := c.actors[actorID]
	return ok
}

func (c *Container) onActorCreate(actor Actor) {
	c.actors[actor.ID()] = actor
	if actor.ActorType() == battle.ActorType_Submarine {
		c.submarines[actor.UserID()] = actor
	}
	actor.Start()
	c.context.Event.EmitSync(event.ActorAdded, actor)
}

func (c *Container) onActorDestroy(rawActor Actor) {
	actor := c.Actor(rawActor.ID())
	if actor == nil {
		return
	}

	delete(c.actors, actor.ID())
	if actor.ActorType() == battle.ActorType_Submarine {
		delete(c.submarines, actor.UserID())
	}
	actor.OnDestroy()
	c.context.Event.EmitSync(event.ActorRemoved, actor)
}
