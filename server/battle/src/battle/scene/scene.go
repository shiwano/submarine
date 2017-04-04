package scene

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Scene represents a battle scene.
type Scene interface {
	Event() *EventEmitter
	Stage() *navmesh.NavMesh
	SightsByTeam() map[navmesh.LayerMask]*sight.Sight
	Now() time.Time
	Actors() ActorSlice
	Actor(actorID int64) (Actor, bool)
	HasActor(actorID int64) bool
}

// FullScene represents a battle scene, includes methods that manages itself.
type FullScene interface {
	Scene
	StartedAt() time.Time
	ElapsedTime() time.Duration
	Start(now time.Time)
	Update(now time.Time)
	Players() PlayerSlice
	UserPlayersByTeam() PlayersByTeam
	SubmarineByPlayerID(userID int64) (Actor, bool)
}

type scene struct {
	event        *EventEmitter
	stage        *navmesh.NavMesh
	sightsByTeam map[navmesh.LayerMask]*sight.Sight
	startedAt    time.Time
	now          time.Time
	container    *container
}

// NewScene creates a battle scene.
func NewScene(stageMesh *navmesh.Mesh, lightMap *sight.LightMap) FullScene {
	scn := &scene{
		event:        NewEventEmitter(),
		stage:        navmesh.New(stageMesh),
		sightsByTeam: make(map[navmesh.LayerMask]*sight.Sight),
		container:    newContainer(),
	}
	for _, layer := range TeamLayers {
		scn.sightsByTeam[layer] = sight.New(lightMap)
	}
	scn.event.AddActorCreateEventListener(scn.onActorCreate)
	scn.event.AddActorDestroyEventListener(scn.onActorDestroy)
	return scn
}

func (scn *scene) Now() time.Time                                   { return scn.now }
func (scn *scene) StartedAt() time.Time                             { return scn.startedAt }
func (scn *scene) ElapsedTime() time.Duration                       { return scn.now.Sub(scn.startedAt) }
func (scn *scene) Event() *EventEmitter                             { return scn.event }
func (scn *scene) Stage() *navmesh.NavMesh                          { return scn.stage }
func (scn *scene) SightsByTeam() map[navmesh.LayerMask]*sight.Sight { return scn.sightsByTeam }

func (scn *scene) Players() PlayerSlice             { return scn.container.players }
func (scn *scene) UserPlayersByTeam() PlayersByTeam { return scn.container.userPlayersByTeam }

func (scn *scene) Start(now time.Time) {
	scn.startedAt = now
	scn.now = now
}

func (scn *scene) Update(now time.Time) {
	scn.now = now
	for _, sight := range scn.sightsByTeam {
		sight.Clear()
	}
	for _, actor := range scn.Actors() {
		if !actor.IsDestroyed() {
			actor.BeforeUpdate()
		}
	}
	for _, actor := range scn.Actors() {
		if !actor.IsDestroyed() {
			actor.Update()
		}
	}
	for _, actor := range scn.Actors() {
		if !actor.IsDestroyed() {
			actor.AfterUpdate()
		}
	}
}

func (scn *scene) SubmarineByPlayerID(userID int64) (Actor, bool) {
	if s, ok := scn.container.submarinesByPlayerID[userID]; ok {
		return s, true
	}
	return nil, false
}

func (scn *scene) Actors() ActorSlice {
	actors := make(ActorSlice, len(scn.container.actors))
	copy(actors, scn.container.actors)
	return actors
}

func (scn *scene) Actor(actorID int64) (Actor, bool) {
	if a, ok := scn.container.actorsByID[actorID]; ok {
		return a, true
	}
	return nil, false
}

func (scn *scene) HasActor(actorID int64) bool {
	_, ok := scn.container.actorsByID[actorID]
	return ok
}

func (scn *scene) onActorCreate(actor Actor) {
	scn.container.addActor(actor)
	actor.Start()
	scn.event.EmitActorAddEvent(actor)
}

func (scn *scene) onActorDestroy(actor Actor) {
	removedActor := scn.container.removeActor(actor)
	if removedActor != nil {
		removedActor.OnDestroy()
		scn.event.EmitActorRemoveEvent(removedActor)
	}
}
