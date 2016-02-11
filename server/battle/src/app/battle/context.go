package battle

import (
	"github.com/chuckpreslar/emission"
	"time"
)

// Context represents a battle context.
type Context struct {
	now       time.Time
	container *ActorContainer
	event     *emission.Emitter
}

func newContext() *Context {
	context := new(Context)
	context.event = emission.NewEmitter()
	context.container = newActorContainer(context)
	return context
}
