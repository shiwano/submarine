package event

import (
	"github.com/chuckpreslar/emission"
)

// Emitter emits battle events.
type Emitter struct {
	emitter *emission.Emitter
}

// New creates a Emitter.
func New() *Emitter {
	return &Emitter{
		emitter: emission.NewEmitter(),
	}
}

// On adds a event listener.
func (e *Emitter) On(event Type, listener interface{}) *Emitter {
	e.emitter.On(event, listener)
	return e
}

// Off removes a event listener.
func (e *Emitter) Off(event Type, listener interface{}) *Emitter {
	e.emitter.Off(event, listener)
	return e
}

// Once adds a event listener which invokes only once.
func (e *Emitter) Once(event Type, listener interface{}) *Emitter {
	e.emitter.Once(event, listener)
	return e
}

// Emit emits the specified event synchronously.
func (e *Emitter) Emit(event Type, arguments ...interface{}) *Emitter {
	e.emitter.EmitSync(event, arguments...)
	return e
}
