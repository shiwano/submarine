package atomicbool

import (
	"sync"
)

// T represents an atomic bool.
type T struct {
	mutex *sync.Mutex
	value bool
}

// New creates an atomic bool.
func New(value bool) *T {
	return &T{
		mutex: new(sync.Mutex),
		value: value,
	}
}

// Value returns the value.
func (t *T) Value() bool {
	t.mutex.Lock()
	value := t.value
	t.mutex.Unlock()
	return value
}

// Set the value.
func (t *T) Set(value bool) {
	t.mutex.Lock()
	t.value = value
	t.mutex.Unlock()
}

// Toggle the value.
func (t *T) Toggle() {
	t.mutex.Lock()
	t.value = !t.value
	t.mutex.Unlock()
}
