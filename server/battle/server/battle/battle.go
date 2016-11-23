package battle

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/actor"
	"github.com/shiwano/submarine/server/battle/server/battle/ai"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/debug"

	"github.com/tevino/abool"
	"github.com/ungerik/go3d/float64/vec2"
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
func New(timeLimit time.Duration, stageMesh *navmesh.Mesh, lightMap *sight.LightMap) *Battle {
	ctx := context.NewContext(stageMesh, lightMap)
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

// Start starts the battle that is startable.
func (b *Battle) Start() bool {
	// TODO: Relevant users counting.
	if !b.isStarted && len(b.ctx.Players()) > 0 {
		b.isStarted = true
		go b.run()
		return true
	}
	return false
}

// Close closes the battle that is running.
func (b *Battle) Close() {
	if b.isStarted && b.isFighting.IsSet() {
		b.closeCh <- struct{}{}
	}
}

// EnterUser enters an user to the battle.
func (b *Battle) EnterUser(userID int64) {
	if !b.isStarted {
		if _, ok := b.ctx.SubmarineByPlayerID(userID); !ok {
			index := len(b.ctx.Players())
			startPos := b.getStartPosition(index)
			teamLayer := context.GetTeamLayer(index + 1)
			user := context.NewPlayer(userID, true, teamLayer, startPos)
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

// EnterBot enters a bot to the battle.
func (b *Battle) EnterBot(bot *api.Bot) bool {
	if !b.isStarted {
		index := len(b.ctx.Players())
		startPos := b.getStartPosition(index)
		teamLayer := context.GetTeamLayer(index + 1)
		player := context.NewPlayer(bot.Id, false, teamLayer, startPos)
		player.AI = ai.NewSimpleAI(b.ctx)
		actor.NewSubmarine(b.ctx, player)
		return true
	}
	return false
}

// LeaveBot leaves a bot from the battle.
func (b *Battle) LeaveBot(bot *api.Bot) bool {
	if !b.isStarted {
		if s, ok := b.ctx.SubmarineByPlayerID(bot.Id); ok {
			s.Destroy()
		}
		return true
	}
	return false
}

func (b *Battle) run() {
	b.start()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()
loop:
	for {
		select {
		case now := <-ticker.C:
			if b.update(now) {
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
	b.Gateway.outputStart(b.ctx.Players(), b.ctx.StartedAt)
	for _, actor := range b.ctx.Actors() {
		b.Gateway.outputActor(b.ctx.UserPlayersByTeam(), actor)
	}
	b.ctx.Event.AddActorAddEventListener(b.onActorAdd)
	b.ctx.Event.AddActorMoveEventListener(b.onActorMove)
	b.ctx.Event.AddActorChangeVisibilityEventListener(b.onActorChangeVisibility)
	b.ctx.Event.AddActorDestroyEventListener(b.onActorDestroy)

	if debug.Debug {
		debug.Debugger.Update(b.ctx.Stage)
	}
}

func (b *Battle) update(now time.Time) bool {
	b.ctx.Update(now)
	if debug.Debug {
		debug.Debugger.Update(b.ctx.Stage)
	}
	return b.judge.isBattleFinished()
}

func (b *Battle) finish() {
	b.isFighting.SetTo(false)
	if winner := b.judge.winner(); winner != nil {
		b.Gateway.outputFinish(&winner.ID, b.ctx.Now)
	} else {
		b.Gateway.outputFinish(nil, b.ctx.Now)
	}
	if debug.Debug {
		debug.Debugger.Update(nil)
	}
}

func (b *Battle) reenterUser(userID int64) {
	if s, ok := b.ctx.SubmarineByPlayerID(userID); ok {
		players := context.PlayerSlice{s.Player()}
		b.Gateway.outputStart(players, b.ctx.StartedAt)
		for _, actor := range b.ctx.Actors() {
			b.Gateway.outputActor(players.GroupByTeam(), actor)
		}
	}
}

func (b *Battle) leaveUser(userID int64) {
	if s, ok := b.ctx.SubmarineByPlayerID(userID); ok {
		s.Event().EmitUserLeaveEvent()
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
	s, ok := b.ctx.SubmarineByPlayerID(input.userID)
	if !ok {
		return
	}
	switch m := input.message.(type) {
	case *battleAPI.AccelerationRequestObject:
		s.Event().EmitAccelerationRequestEvent(m)
	case *battleAPI.BrakeRequestObject:
		s.Event().EmitBrakeRequestEvent(m)
	case *battleAPI.TurnRequestObject:
		s.Event().EmitTurnRequestEvent(m)
	case *battleAPI.TorpedoRequestObject:
		s.Event().EmitTorpedoRequestEvent(m)
	case *battleAPI.PingerRequestObject:
		s.Event().EmitPingerRequestEvent(m)
	}
}

func (b *Battle) onActorAdd(actor context.Actor) {
	b.Gateway.outputActor(b.ctx.UserPlayersByTeam(), actor)
}

func (b *Battle) onActorMove(actor context.Actor) {
	b.Gateway.outputMovement(actor)
}

func (b *Battle) onActorChangeVisibility(actor context.Actor, teamLayer navmesh.LayerMask) {
	b.Gateway.outputVisibility(b.ctx.UserPlayersByTeam(), actor, teamLayer)
}

func (b *Battle) onActorDestroy(actor context.Actor) {
	b.Gateway.outputDestruction(actor)
}
