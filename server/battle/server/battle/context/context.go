package context

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Context represents a battle context.
type Context interface {
	Event() *EventEmitter
	Stage() *navmesh.NavMesh
	SightsByTeam() map[navmesh.LayerMask]*sight.Sight
	Now() time.Time
	StartedAt() time.Time
	ElapsedTime() time.Duration
	Actors() ActorSlice
	Actor(actorID int64) (Actor, bool)
	HasActor(actorID int64) bool
}

// FullContext represents a battle context, includes methods that manages itself.
type FullContext interface {
	Context
	Start(now time.Time)
	Update(now time.Time)
	Players() PlayerSlice
	UserPlayersByTeam() PlayersByTeam
	SubmarineByPlayerID(userID int64) (Actor, bool)
}

type context struct {
	event        *EventEmitter
	stage        *navmesh.NavMesh
	sightsByTeam map[navmesh.LayerMask]*sight.Sight
	startedAt    time.Time
	now          time.Time
	container    *container
}

// NewContext creates a battle context.
func NewContext(stageMesh *navmesh.Mesh, lightMap *sight.LightMap) FullContext {
	c := &context{
		event:        NewEventEmitter(),
		stage:        navmesh.New(stageMesh),
		sightsByTeam: make(map[navmesh.LayerMask]*sight.Sight),
		container:    newContainer(),
	}
	for _, layer := range TeamLayers {
		c.sightsByTeam[layer] = sight.New(lightMap)
	}
	c.event.AddActorCreateEventListener(c.onActorCreate)
	c.event.AddActorDestroyEventListener(c.onActorDestroy)
	return c
}

func (c *context) Now() time.Time                                   { return c.now }
func (c *context) StartedAt() time.Time                             { return c.startedAt }
func (c *context) ElapsedTime() time.Duration                       { return c.now.Sub(c.startedAt) }
func (c *context) Event() *EventEmitter                             { return c.event }
func (c *context) Stage() *navmesh.NavMesh                          { return c.stage }
func (c *context) SightsByTeam() map[navmesh.LayerMask]*sight.Sight { return c.sightsByTeam }

func (c *context) Players() PlayerSlice             { return c.container.players }
func (c *context) UserPlayersByTeam() PlayersByTeam { return c.container.userPlayersByTeam }

func (c *context) Start(now time.Time) {
	c.startedAt = now
	c.now = now
}

func (c *context) Update(now time.Time) {
	c.now = now
	for _, sight := range c.sightsByTeam {
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

func (c *context) SubmarineByPlayerID(userID int64) (Actor, bool) {
	if s, ok := c.container.submarinesByPlayerID[userID]; ok {
		return s, true
	}
	return nil, false
}

func (c *context) Actors() ActorSlice {
	actors := make(ActorSlice, len(c.container.actors))
	copy(actors, c.container.actors)
	return actors
}

func (c *context) Actor(actorID int64) (Actor, bool) {
	if a, ok := c.container.actorsByID[actorID]; ok {
		return a, true
	}
	return nil, false
}

func (c *context) HasActor(actorID int64) bool {
	_, ok := c.container.actorsByID[actorID]
	return ok
}

func (c *context) onActorCreate(actor Actor) {
	c.container.addActor(actor)
	actor.Start()
	c.event.EmitActorAddEvent(actor)
}

func (c *context) onActorDestroy(actor Actor) {
	removedActor := c.container.removeActor(actor)
	if removedActor != nil {
		removedActor.OnDestroy()
		c.event.EmitActorRemoveEvent(removedActor)
	}
}
