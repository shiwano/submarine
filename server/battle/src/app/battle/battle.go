package battle

import (
	"app/battle/actor"
	"app/battle/context"
	"app/battle/event"
	"app/logger"
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
	return b
}

// CreateSubmarineUnlessExists creates the user's submarine unless it exists.
func (b *Battle) CreateSubmarineUnlessExists(userID int64) {
	if s := b.context.SubmarineByUserID(userID); s == nil {
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
	b.context.Event.On(event.ActorAdd, b.onActorAdd)
	b.context.Event.On(event.ActorMove, b.onActorMove)

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
	for _, actor := range b.context.Actors() {
		actor.Update()
	}
	return b.context.Now.Before(b.startedAt.Add(b.timeLimit))
}

func (b *Battle) onInputReceive(input *gatewayInput) {
	s := b.context.SubmarineByUserID(input.userID)
	if s == nil {
		return
	}

	switch m := input.message.(type) {
	case *battle.AccelerationRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) accelerates", s.UserID(), s.ID())
		s.Event().Emit(event.AccelerationRequest, m)
	case *battle.BrakeRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) brakes", s.UserID(), s.ID())
		s.Event().Emit(event.BrakeRequest, m)
	case *battle.TurnRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) turns to %v", s.UserID(), s.ID(), m.Direction)
		s.Event().Emit(event.TurnRequest, m)
	case *battle.TorpedoRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) shoots a torpedo", s.UserID(), s.ID())
		s.Event().Emit(event.TorpedoRequest, m)
	case *battle.PingerRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) use pinger", s.UserID(), s.ID())
		s.Event().Emit(event.PingerRequest, m)
	}
}

func (b *Battle) onActorAdd(actor context.Actor) {
	b.Gateway.outputActor(actor)
}

func (b *Battle) onActorMove(actor context.Actor) {
	b.Gateway.outputMovement(actor)
}
