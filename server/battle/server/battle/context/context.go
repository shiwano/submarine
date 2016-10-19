package context

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
	"github.com/shiwano/submarine/server/battle/server/battle/event"
)

// Context represents a battle context.
type Context struct {
	CreatedAt    time.Time
	StartedAt    time.Time
	Now          time.Time
	Event        *event.Emitter
	Stage        *navmesh.NavMesh
	SightsByTeam map[navmesh.LayerMask]*sight.Sight
	container    *container
}

// NewContext creates a battle context.
func NewContext(stageMesh *navmesh.Mesh, lightMap *sight.LightMap) *Context {
	c := &Context{
		CreatedAt:    time.Now(),
		Event:        event.New(),
		Stage:        navmesh.New(stageMesh),
		SightsByTeam: make(map[navmesh.LayerMask]*sight.Sight),
		container:    newContainer(),
	}
	for _, layer := range TeamLayers {
		c.SightsByTeam[layer] = sight.New(lightMap)
	}
	c.Event.On(event.ActorCreate, c.onActorCreate)
	c.Event.On(event.ActorDestroy, c.onActorDestroy)
	return c
}

// Update the battle.
func (c *Context) Update(now time.Time) {
	c.Now = now
	for _, sight := range c.SightsByTeam {
		sight.Clear()
	}
	for _, actor := range c.Actors() {
		if !actor.IsDestroyed() {
			actor.BeforeUpdate()
		}
	}
	for _, actor := range c.Actors() {
		if !actor.IsDestroyed() {
			actor.Update()
		}
	}
	for _, actor := range c.Actors() {
		if !actor.IsDestroyed() {
			actor.AfterUpdate()
		}
	}
}

// ElapsedTime returns the elapsed time since start of battle.
func (c *Context) ElapsedTime() time.Duration {
	return c.Now.Sub(c.StartedAt)
}

// SubmarineByPlayerID returns the submarine which has the given player id.
func (c *Context) SubmarineByPlayerID(userID int64) Actor {
	return c.container.submarinesByPlayerID[userID]
}

// Actors returns all actors.
func (c *Context) Actors() ActorSlice {
	actors := make(ActorSlice, len(c.container.actors))
	copy(actors, c.container.actors)
	return actors
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

// Players returns players in the battle.
func (c *Context) Players() PlayerSlice {
	return c.container.players
}

// UserPlayersByTeam returns user's players by team layer of the battle.
func (c *Context) UserPlayersByTeam() PlayersByTeam {
	return c.container.userPlayersByTeam
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
