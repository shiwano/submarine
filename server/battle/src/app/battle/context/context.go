package context

import (
	"github.com/chuckpreslar/emission"
	"time"
)

// Context represents a battle context.
type Context struct {
	currentActorID int64
	Now            time.Time
	Event          *emission.Emitter
	Container      *Container
}

// NewContext creates a contest.
func NewContext() *Context {
	c := &Context{
		Event: emission.NewEmitter(),
	}
	c.Container = newContainer(c)
	return c
}

// NextActorID returns the next unique actor id.
func (c *Context) NextActorID() int64 {
	nextActorID := c.currentActorID
	c.currentActorID++
	return nextActorID
}

// UserIDs returns user ids in battle.
func (c *Context) UserIDs() []int64 {
	keys := make([]int64, len(c.Container.submarines))
	i := 0
	for k := range c.Container.submarines {
		keys[i] = k
		i++
	}
	return keys
}
