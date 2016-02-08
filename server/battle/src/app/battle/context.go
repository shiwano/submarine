package battle

// Context represents a battle context.
type Context struct {
	now       int64
	container *ActorContainer
}

func newContext() *Context {
	return &Context{
		container: &ActorContainer{},
	}
}
