package battle

import (
	"app/currentmillis"
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway    *Gateway
	context    *Context
	createdAt  int64
	startedAt  int64
	timeLimit  int64
	hasStarted bool
}

// New creates a new battle.
func New(timeLimit int64) *Battle {
	return &Battle{
		Gateway:   newGateway(),
		context:   newContext(),
		createdAt: currentmillis.Now(),
		timeLimit: timeLimit,
	}
}

// Start the battle.
func (b *Battle) Start() {
	if !b.hasStarted {
		b.hasStarted = true
		go b.run()
	}
}

func (b *Battle) run() {
	ticker := time.Tick(time.Second / 30)
	b.startedAt = currentmillis.Now()
	b.Gateway.start(b.startedAt)

loop:
	for {
		select {
		case now := <-ticker:
			b.context.now = currentmillis.Milliseconds(now)
			if b.context.now >= b.startedAt+b.timeLimit {
				break loop
			}
		case <-b.Gateway.Close:
			break loop
		}
	}

	b.Gateway.finish(true, b.context.now)
}
