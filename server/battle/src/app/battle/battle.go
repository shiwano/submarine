package battle

import (
	"app/battle/actor"
	"app/battle/context"
	"app/battle/event"
	"app/logger"
	battleAPI "app/typhenapi/type/submarine/battle"
	"github.com/tevino/abool"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway       *Gateway
	ctx           *context.Context
	judge         *judge
	isStarted     bool
	isFighting    *abool.AtomicBool
	reenterUserCh chan int64
	leaveUserCh   chan int64
	closeCh       chan struct{}
}

// New creates a new battle.
func New(timeLimit time.Duration, stageMesh *navmesh.Mesh) *Battle {
	ctx := context.NewContext(stageMesh)
	return &Battle{
		Gateway:       newGateway(),
		ctx:           ctx,
		judge:         newJudge(ctx, timeLimit),
		isFighting:    abool.New(),
		reenterUserCh: make(chan int64, 4),
		leaveUserCh:   make(chan int64, 4),
		closeCh:       make(chan struct{}, 1),
	}
}

// StartIfPossible starts the battle that is startable.
func (b *Battle) StartIfPossible() bool {
	// TODO: Relevant users counting.
	if !b.isStarted && len(b.ctx.Users()) > 0 {
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
		if s := b.ctx.SubmarineByUserID(userID); s == nil {
			index := len(b.ctx.Users())
			startPos := b.getStartPosition(index)
			teamLayer := context.GetTeamLayer(index + 1)
			user := context.NewUser(userID, teamLayer, startPos)
			actor.NewSubmarine(b.ctx, user)
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
			b.update(now)
			if b.judge.isBattleFinished() {
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
	b.ctx.StartedAt = time.Now()
	b.Gateway.outputStart(nil, b.ctx.StartedAt)
	for _, actor := range b.ctx.Actors() {
		b.Gateway.outputActor(nil, actor)
	}
	b.ctx.Event.On(event.ActorAdd, b.onActorAdd)
	b.ctx.Event.On(event.ActorMove, b.onActorMove)
	b.ctx.Event.On(event.ActorDestroy, b.onActorDestroy)
}

func (b *Battle) update(now time.Time) {
	b.ctx.Now = now
	for _, actor := range b.ctx.Actors() {
		if !actor.IsDestroyed() {
			actor.BeforeUpdate()
		}
		if !actor.IsDestroyed() {
			actor.Update()
		}
	}
}

func (b *Battle) finish() {
	b.isFighting.SetTo(false)
	if winner := b.judge.winner(); winner != nil {
		b.Gateway.outputFinish(&winner.ID, b.ctx.Now)
	} else {
		b.Gateway.outputFinish(nil, b.ctx.Now)
	}
}

func (b *Battle) reenterUser(userID int64) {
	userIDs := []int64{userID}
	b.Gateway.outputStart(userIDs, b.ctx.StartedAt)
	for _, actor := range b.ctx.Actors() {
		b.Gateway.outputActor(userIDs, actor)
	}
}

func (b *Battle) leaveUser(userID int64) {
	s := b.ctx.SubmarineByUserID(userID)
	if s != nil {
		s.Event().Emit(event.UserLeave)
	}
}

func (b *Battle) getStartPosition(index int) *vec2.T {
	switch index {
	case 0:
		return &vec2.T{125, 125}
	case 1:
		return &vec2.T{-125, -125}
	case 2:
		return &vec2.T{125, -125}
	default:
		return &vec2.T{-125, 125}
	}
}

func (b *Battle) onInputReceive(input *gatewayInput) {
	s := b.ctx.SubmarineByUserID(input.userID)
	if s == nil {
		return
	}
	switch m := input.message.(type) {
	case *battleAPI.AccelerationRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) accelerates", s.User().ID, s.ID())
		s.Event().Emit(event.AccelerationRequest, m)
	case *battleAPI.BrakeRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) brakes", s.User().ID, s.ID())
		s.Event().Emit(event.BrakeRequest, m)
	case *battleAPI.TurnRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) turns to %v", s.User().ID, s.ID(), m.Direction)
		s.Event().Emit(event.TurnRequest, m)
	case *battleAPI.TorpedoRequestObject:
		logger.Log.Debugf("User(%v)'s submarine(%v) shoots a torpedo", s.User().ID, s.ID())
		s.Event().Emit(event.TorpedoRequest, m)
	case *battleAPI.PingerRequestObject:
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
