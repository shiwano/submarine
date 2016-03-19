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
	Gateway       *Gateway
	context       *context.Context
	createdAt     time.Time
	startedAt     time.Time
	timeLimit     time.Duration
	IsStarted     bool
	reenterUserCh chan int64
	closeCh       chan struct{}
}

// New creates a new battle.
func New(timeLimit time.Duration) *Battle {
	battleContext := context.NewContext()
	b := &Battle{
		Gateway:       newGateway(),
		context:       battleContext,
		createdAt:     time.Now(),
		timeLimit:     timeLimit,
		reenterUserCh: make(chan int64, 4),
		closeCh:       make(chan struct{}, 1),
	}
	return b
}

// EnterUser an user to the battle.
func (b *Battle) EnterUser(userID int64) {
	if b.IsStarted {
		b.reenterUserCh <- userID
	} else {
		if s := b.context.SubmarineByUserID(userID); s == nil {
			actor.NewSubmarine(b.context, userID)
		}
	}
}

// Start the battle.
func (b *Battle) Start() {
	if !b.IsStarted {
		b.IsStarted = true
		go b.run()
	}
}

// Close the battle.
func (b *Battle) Close() {
	if b.IsStarted {
		b.closeCh <- struct{}{}
	}
}

func (b *Battle) run() {
	b.start()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()
loop:
	for {
		select {
		case now := <-ticker.C:
			if !b.update(now) {
				break loop
			}
		case input := <-b.Gateway.input:
			b.onInputReceive(input)
		case userID := <-b.reenterUserCh:
			b.reenterUser(userID)
		case <-b.closeCh:
			break loop
		}
	}
	b.finish()
}

func (b *Battle) start() {
	b.startedAt = time.Now()
	b.Gateway.outputStart(nil, b.startedAt)
	for _, actor := range b.context.Actors() {
		b.Gateway.outputActor(nil, actor)
	}
	b.context.Event.On(event.ActorAdd, b.onActorAdd)
	b.context.Event.On(event.ActorMove, b.onActorMove)
}

func (b *Battle) update(now time.Time) bool {
	b.context.Now = now
	for _, actor := range b.context.Actors() {
		actor.Update()
	}
	return b.context.Now.Before(b.startedAt.Add(b.timeLimit))
}

func (b *Battle) finish() {
	// TODO: winnerUserID is temporary value.
	b.Gateway.outputFinish(b.context.UserIDs()[0], b.context.Now)
}

func (b *Battle) reenterUser(userID int64) {
	userIDs := []int64{userID}
	b.Gateway.outputStart(userIDs, b.startedAt)
	for _, actor := range b.context.Actors() {
		b.Gateway.outputActor(userIDs, actor)
	}
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
	b.Gateway.outputActor(nil, actor)
}

func (b *Battle) onActorMove(actor context.Actor) {
	b.Gateway.outputMovement(actor)
}
