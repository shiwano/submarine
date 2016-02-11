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
	return &Context{
		container: newActorContainer(),
		event:     emission.NewEmitter(),
	}
}
