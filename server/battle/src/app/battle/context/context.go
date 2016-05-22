package context

import (
	"app/battle/event"
	"lib/navmesh"
	"time"
)

// Context represents a battle context.
type Context struct {
	Now       time.Time
	Event     *event.Emitter
	Stage     *navmesh.NavMesh
	container *container
}

// NewContext creates a contest.
func NewContext(stageMesh *navmesh.Mesh) *Context {
	c := &Context{
		Event:     event.New(),
		Stage:     navmesh.New(stageMesh),
		container: newContainer(),
	}
	c.Event.On(event.ActorCreate, c.onActorCreate)
	c.Event.On(event.ActorDestroy, c.onActorDestroy)
	return c
}

// NextActorID returns the next unique actor id.
func (c *Context) NextActorID() int64 {
	return c.container.nextActorID()
}

// SubmarineByUserID returns the submarine which has the given actor id.
func (c *Context) SubmarineByUserID(userID int64) Actor {
	return c.container.submarinesByUserID[userID]
}

// Actors returns all actors.
func (c *Context) Actors() []Actor {
	return c.container.actors
}

// Actor returns the actor that has the actor id.
func (c *Context) Actor(actorID int64) Actor {
	return c.container.actorsByID[actorID]
}

// HasActor determines whether the specified actor exists.
func (c *Context) HasActor(actorID int64) bool {
	_, ok := c.container.actorsByID[actorID]
	return ok
}

// UserIDs returns user ids in battle.
func (c *Context) UserIDs() []int64 {
	keys := make([]int64, len(c.container.submarinesByUserID))
	i := 0
	for k := range c.container.submarinesByUserID {
		keys[i] = k
		i++
	}
	return keys
}

func (c *Context) onActorCreate(actor Actor) {
	c.container.addActor(actor)
	actor.Start()
	c.Event.Emit(event.ActorAdd, actor)
}

func (c *Context) onActorDestroy(actor Actor) {
	removedActor := c.container.removeActor(actor)
	if removedActor != nil {
		removedActor.OnDestroy()
		c.Event.Emit(event.ActorRemove, removedActor)
	}
}
