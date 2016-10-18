package context

import (
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/event"
)

var lastCreateActorID int64

type actor struct {
	id          int64
	player      *Player
	actorType   battleAPI.ActorType
	ctx         *Context
	event       *event.Emitter
	isDestroyed bool

	isCalledStart     bool
	isCalledOnDestroy bool
}

func newSubmarine(ctx *Context, isUser bool) *actor {
	lastCreateActorID++
	id := lastCreateActorID
	player := &Player{ID: id * 100, IsUser: isUser, StartPosition: &vec2.Zero}
	a := &actor{
		id:        id,
		player:    player,
		actorType: battleAPI.ActorType_Submarine,
		ctx:       ctx,
		event:     event.New(),
	}
	a.ctx.Event.Emit(event.ActorCreate, a)
	return a
}

func (a *actor) ID() int64                 { return a.id }
func (a *actor) Player() *Player           { return a.player }
func (a *actor) Type() battleAPI.ActorType { return a.actorType }
func (a *actor) Event() *event.Emitter     { return a.event }

func (a *actor) IsDestroyed() bool                    { return a.isDestroyed }
func (a *actor) Movement() *battleAPI.Movement        { panic("not implemented yet.") }
func (a *actor) Position() *vec2.T                    { return &vec2.Zero }
func (a *actor) Direction() float64                   { return 0 }
func (a *actor) IsAccelerating() bool                 { return false }
func (a *actor) IsVisibleFrom(navmesh.LayerMask) bool { return true }

func (a *actor) Destroy() {
	a.isDestroyed = true
	a.ctx.Event.Emit(event.ActorDestroy, a)
}

func (a *actor) Start()        { a.isCalledStart = true }
func (a *actor) BeforeUpdate() {}
func (a *actor) Update()       {}
func (a *actor) OnDestroy()    { a.isCalledOnDestroy = true }
