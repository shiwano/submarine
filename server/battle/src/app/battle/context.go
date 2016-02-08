package battle

import (
	"time"
)

// Context represents a battle context.
type Context struct {
	now       time.Time
	container *ActorContainer
}

func newContext() *Context {
	return &Context{
		container: newActorContainer(),
	}
}
