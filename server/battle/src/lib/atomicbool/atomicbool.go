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
	defer t.mutex.Unlock()
	return t.value
}

// Set the value.
func (t *T) Set(value bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.value = value
}

// Toggle the value.
func (t *T) Toggle() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.value = !t.value
}
