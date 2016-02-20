package battle

import (
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway   *Gateway
	context   *Context
	createdAt time.Time
	startedAt time.Time
	timeLimit time.Duration
	IsStarted bool
	close     chan struct{}
}

// New creates a new battle.
func New(timeLimit time.Duration) *Battle {
	battle := &Battle{
		Gateway:   newGateway(),
		context:   newContext(),
		createdAt: time.Now(),
		timeLimit: timeLimit,
		close:     make(chan struct{}, 1),
	}

	battle.context.event.On(ActorCreated, func(actor Actor) {
		battle.Gateway.actor(actor)
	})
	return battle
}

// CreateSubmarineUnlessExists creates the user's submarine unless it exists.
func (b *Battle) CreateSubmarineUnlessExists(userID int64) {
	if s := b.context.container.getSubmarineByUserID(userID); s == nil {
		b.context.container.createSubmarine(userID)
	}
}

// Start the battle.
func (b *Battle) Start() {
	if !b.IsStarted {
		b.start()
		go b.run()
	}
}

// Close the battle.
func (b *Battle) Close() {
	if b.IsStarted {
		b.close <- struct{}{}
	}
}

func (b *Battle) run() {
	ticker := time.Tick(time.Second / 30)
	b.Gateway.start(b.startedAt)

loop:
	for {
		select {
		case now := <-ticker:
			if !b.update(now) {
				break loop
			}
		case <-b.close:
			break loop
		}
	}

	// TODO: winnerUserID is temporary value.
	b.Gateway.finish(b.context.userIDs()[0], b.context.now)
}

func (b *Battle) start() {
	b.IsStarted = true
	b.startedAt = time.Now()
}

func (b *Battle) update(now time.Time) bool {
	b.context.now = now
	for _, actor := range b.context.container.actors {
		actor.Update()
	}
	return b.context.now.Before(b.startedAt.Add(b.timeLimit))
}
