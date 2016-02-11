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

func (c *Context) userIDs() []int64 {
	keys := make([]int64, len(c.container.submarines))
	i := 0
	for k := range c.container.submarines {
		keys[i] = k
		i++
	}
	return keys
}
