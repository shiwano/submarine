package atomicbool

import (
	"sync/atomic"
)

// T represents an atomic bool.
type T struct {
	value int32
}

// New creates an atomic bool.
func New(value bool) *T {
	return &T{value: toInt32(value)}
}

func toInt32(value bool) int32 {
	if value {
		return 1
	}
	return 0
}

// Set sets the value.
func (t *T) Set(value bool) {
	atomic.StoreInt32(&(t.value), int32(toInt32(value)))
}

// Value returns the value.
func (t *T) Value() bool {
	return atomic.LoadInt32(&(t.value)) == 1
}
