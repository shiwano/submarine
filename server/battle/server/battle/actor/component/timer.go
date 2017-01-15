package component

import "time"

// Timer executes a function when the specified interval elapsed.
type Timer struct {
	now   time.Time
	items []*TimerItem
}

// NewTimer creates a timer.
func NewTimer(now time.Time) *Timer {
	return &Timer{
		now: now,
	}
}

// Update registered timer items.
func (t *Timer) Update(now time.Time) {
	t.now = now
	if len(t.items) == 0 {
		return
	}

	items := t.items[:0]
	for _, item := range t.items {
		if item.isCanceled {
			continue
		}
		if t.now.After(item.elapsesAt) {
			if item.elapseHandler != nil {
				item.elapseHandler()
			}
			if !item.repeat {
				continue
			}
		}
		items = append(items, item)
	}
	t.items = items
}

// Register a timer item.
func (t *Timer) Register(intervalSeconds float64, handler func()) *TimerItem {
	return t.register(false, intervalSeconds, handler)
}

// RegisterRepeat a repeat timer item.
func (t *Timer) RegisterRepeat(intervalSeconds float64, handler func()) *TimerItem {
	return t.register(true, intervalSeconds, handler)
}

func (t *Timer) register(repeat bool, intervalSeconds float64, handler func()) *TimerItem {
	interval := time.Duration(intervalSeconds * float64(time.Second))
	item := &TimerItem{
		repeat:        repeat,
		elapsesAt:     t.now.Add(interval),
		elapseHandler: handler,
	}
	t.items = append(t.items, item)
	return item
}

// TimerItem represents an registered timer item.
type TimerItem struct {
	repeat        bool
	elapsesAt     time.Time
	elapseHandler func()
	isCanceled    bool
}

// Cancel the timer item.
func (i *TimerItem) Cancel() {
	i.isCanceled = true
}
