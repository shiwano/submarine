package atomicbool

import (
	"sync"
	"testing"
)

func TestAtomicBool(t *testing.T) {
	a := New(true)
	if !a.Value() {
		t.Error("Failed to initialize the value")
		return
	}

	a.Set(false)
	if a.Value() {
		t.Error("Failed to set the value")
		return
	}

	a.Set(true)
	if !a.Value() {
		t.Error("Failed to set the value")
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			a.Toggle()
			wg.Done()
		}(i)
	}
	wg.Wait()
	if !a.Value() {
		t.Error("Failed to toggle the value with thread safe")
		return
	}
}
