package battle

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
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

	battle.context.event.On(event.ActorCreated, func(actor Actor) {
		battle.Gateway.outputActor(actor)
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
	b.Gateway.outputStart(b.startedAt)

loop:
	for {
		select {
		case now := <-ticker:
			if !b.update(now) {
				break loop
			}
		case input := <-b.Gateway.input:
			b.onInputReceive(input)
		case <-b.close:
			break loop
		}
	}

	// TODO: winnerUserID is temporary value.
	b.Gateway.outputFinish(b.context.userIDs()[0], b.context.now)
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

func (b *Battle) onInputReceive(input *Input) {
	submarine := b.context.container.getSubmarineByUserID(input.UserID)
	if submarine == nil {
		return
	}

	switch message := input.Message.(type) {
	case *battle.AccelerationRequestObject:
		submarine.event.EmitSync(event.AccelerationRequested, message)
	case *battle.BrakeRequestObject:
		submarine.event.EmitSync(event.BrakeRequested, message)
	case *battle.TurnRequestObject:
		submarine.event.EmitSync(event.TurnRequested, message)
	case *battle.TorpedoRequestObject:
		submarine.event.EmitSync(event.TorpedoRequested, message)
	case *battle.PingerRequestObject:
		submarine.event.EmitSync(event.PingerRequested, message)
	}
}
