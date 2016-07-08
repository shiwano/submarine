package battle

import (
	"app/battle/actor"
	"app/battle/context"
	"app/battle/event"
	"app/logger"
	"app/typhenapi/type/submarine/battle"
	"github.com/tevino/abool"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway       *Gateway
	context       *context.Context
	createdAt     time.Time
	startedAt     time.Time
	timeLimit     time.Duration
	isStarted     bool
	isFighting    *abool.AtomicBool
	reenterUserCh chan int64
	leaveUserCh   chan int64
	closeCh       chan struct{}
}

// New creates a new battle.
func New(timeLimit time.Duration, stageMesh *navmesh.Mesh) *Battle {
	return &Battle{
		Gateway:       newGateway(),
		context:       context.NewContext(stageMesh),
		createdAt:     time.Now(),
		timeLimit:     timeLimit,
		isFighting:    abool.New(),
		reenterUserCh: make(chan int64, 4),
		leaveUserCh:   make(chan int64, 4),
		closeCh:       make(chan struct{}, 1),
	}
}

// StartIfPossible starts the battle that is startable.
func (b *Battle) StartIfPossible() bool {
	// TODO: Relevant users counting.
	if !b.isStarted && len(b.context.Users()) > 0 {
		b.isStarted = true
		go b.run()
		return true
	}
	return false
}

// CloseIfPossible closes the battle that is running.
func (b *Battle) CloseIfPossible() {
	if b.isStarted && b.isFighting.IsSet() {
		b.closeCh <- struct{}{}
	}
}

// EnterUser enters an user to the battle.
func (b *Battle) EnterUser(userID int64) {
	if !b.isStarted {
		if s := b.context.SubmarineByUserID(userID); s == nil {
			index := len(b.context.Users())
			startPos := &vec2.T{-20 * float64(index), 20 * float64(index)}
			teamLayer := context.GetTeamLayer(index + 1)
			user := context.NewUser(userID, teamLayer, startPos)
			actor.NewSubmarine(b.context, user)
		}
	} else if b.isFighting.IsSet() {
		b.reenterUserCh <- userID
	}
}

// LeaveUser leaves an user from the battle.
func (b *Battle) LeaveUser(userID int64) {
	if b.isFighting.IsSet() {
		b.leaveUserCh <- userID
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
		case userID := <-b.leaveUserCh:
			b.leaveUser(userID)
		case <-b.closeCh:
			break loop
		}
	}
	b.finish()
}

func (b *Battle) start() {
	b.isFighting.SetTo(true)
	b.startedAt = time.Now()
	b.Gateway.outputStart(nil, b.startedAt)
	for _, actor := range b.context.Actors() {
		b.Gateway.outputActor(nil, actor)
	}
	b.context.Event.On(event.ActorAdd, b.onActorAdd)
	b.context.Event.On(event.ActorMove, b.onActorMove)
	b.context.Event.On(event.ActorDestroy, b.onActorDestroy)
}

func (b *Battle) update(now time.Time) bool {
	b.context.Now = now
	for _, actor := range b.context.Actors() {
		if !actor.IsDestroyed() {
			actor.BeforeUpdate()
		}
		if !actor.IsDestroyed() {
			actor.Update()
		}
	}
	return b.context.Now.Before(b.startedAt.Add(b.timeLimit))
}

func (b *Battle) finish() {
	b.isFighting.SetTo(false)
	// TODO: winnerUserID is temporary value.
	b.Gateway.outputFinish(b.context.Users()[0].ID, b.context.Now)
}

func (b *Battle) reenterUser(userID int64) {
	userIDs := []int64{userID}
	b.Gateway.outputStart(userIDs, b.startedAt)
	for _, actor := range b.context.Actors() {
		b.Gateway.outputActor(userIDs, actor)
	}
}

func (b *Battle) leaveUser(userID int64) {
	s := b.context.SubmarineByUserID(userID)
	if s != nil {
		s.Event().Emit(event.UserLeave)
	}
}

func (b *Battle) onInputReceive(input *gatewayInput) {
	s := b.context.SubmarineByUserID(input.userID)
	if s == nil {
		return
	}
	switch m := input.message.(type) {
	case *battle.AccelerationRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) accelerates", s.User().ID, s.ID())
		s.Event().Emit(event.AccelerationRequest, m)
	case *battle.BrakeRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) brakes", s.User().ID, s.ID())
		s.Event().Emit(event.BrakeRequest, m)
	case *battle.TurnRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) turns to %v", s.User().ID, s.ID(), m.Direction)
		s.Event().Emit(event.TurnRequest, m)
	case *battle.TorpedoRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) shoots a torpedo", s.User().ID, s.ID())
		s.Event().Emit(event.TorpedoRequest, m)
	case *battle.PingerRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) use pinger", s.User().ID, s.ID())
		s.Event().Emit(event.PingerRequest, m)
	}
}

func (b *Battle) onActorAdd(actor context.Actor) {
	b.Gateway.outputActor(nil, actor)
}

func (b *Battle) onActorMove(actor context.Actor) {
	b.Gateway.outputMovement(actor)
}

func (b *Battle) onActorDestroy(actor context.Actor) {
	b.Gateway.outputDestruction(actor)
}
