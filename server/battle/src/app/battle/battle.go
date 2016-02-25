package battle

import (
	"app/battle/actor"
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway   *Gateway
	context   *context.Context
	createdAt time.Time
	startedAt time.Time
	timeLimit time.Duration
	IsStarted bool
	close     chan struct{}
}

// New creates a new battle.
func New(timeLimit time.Duration) *Battle {
	battleContext := context.NewContext()
	b := &Battle{
		Gateway:   newGateway(),
		context:   battleContext,
		createdAt: time.Now(),
		timeLimit: timeLimit,
		close:     make(chan struct{}, 1),
	}
	b.context.Event.On(event.ActorAdd, b.onActorAdd)
	return b
}

// CreateSubmarineUnlessExists creates the user's submarine unless it exists.
func (b *Battle) CreateSubmarineUnlessExists(userID int64) {
	if s := b.context.Container.SubmarineByUserID(userID); s == nil {
		actor.NewSubmarine(b.context, userID)
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
	b.Gateway.outputFinish(b.context.UserIDs()[0], b.context.Now)
}

func (b *Battle) start() {
	b.IsStarted = true
	b.startedAt = time.Now()
}

func (b *Battle) update(now time.Time) bool {
	b.context.Now = now
	b.context.Container.UpdateActors()
	return b.context.Now.Before(b.startedAt.Add(b.timeLimit))
}

func (b *Battle) onInputReceive(input *Input) {
	submarine := b.context.Container.SubmarineByUserID(input.UserID)
	if submarine == nil {
		return
	}

	switch message := input.Message.(type) {
	case *battle.AccelerationRequestObject:
		submarine.Event().Emit(event.AccelerationRequest, message)
	case *battle.BrakeRequestObject:
		submarine.Event().Emit(event.BrakeRequest, message)
	case *battle.TurnRequestObject:
		submarine.Event().Emit(event.TurnRequest, message)
	case *battle.TorpedoRequestObject:
		submarine.Event().Emit(event.TorpedoRequest, message)
	case *battle.PingerRequestObject:
		submarine.Event().Emit(event.PingerRequest, message)
	}
}

func (b *Battle) onActorAdd(actor context.Actor) {
	b.Gateway.outputActor(actor)
}
