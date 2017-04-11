package battle

import (
	"context"
	"time"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
	api "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/actor"
	"github.com/shiwano/submarine/server/battle/src/battle/ai"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
	"github.com/shiwano/submarine/server/battle/src/debug"

	"github.com/ungerik/go3d/float64/vec2"
)

// Battle represents a battle.
type Battle struct {
	Gateway *Gateway
	scene   scene.FullScene
	judge   *judge

	userEntered    chan int64
	userLeft       chan int64
	botEntered     chan *api.Bot
	botLeft        chan *api.Bot
	startRequested chan struct{}
	started        chan struct{}
	finished       chan struct{}
	closed         chan struct{}
}

// New creates a new battle.
func New(ctx context.Context, timeLimit time.Duration, stageMesh *navmesh.Mesh, lightMap *sight.LightMap) *Battle {
	scn := scene.NewScene(stageMesh, lightMap)
	b := &Battle{
		Gateway: newGateway(),
		scene:   scn,
		judge:   newJudge(scn, timeLimit),

		userEntered:    make(chan int64),
		userLeft:       make(chan int64),
		botEntered:     make(chan *api.Bot),
		botLeft:        make(chan *api.Bot),
		startRequested: make(chan struct{}),
		started:        make(chan struct{}),
		finished:       make(chan struct{}),
		closed:         make(chan struct{}),
	}
	go b.run(ctx)
	return b
}

// Finished returns a channel that receives a value when the battle finished.
func (b *Battle) Finished() <-chan struct{} {
	return b.finished
}

// Start starts the battle.
func (b *Battle) Start() bool {
	select {
	case <-b.started:
	case <-b.closed:
		return false
	case b.startRequested <- struct{}{}:
	}
	return true
}

// EnterUser enters an user to the battle.
func (b *Battle) EnterUser(userID int64) {
	select {
	case <-b.closed:
	case b.userEntered <- userID:
	}
}

// LeaveUser leaves an user from the battle.
func (b *Battle) LeaveUser(userID int64) {
	select {
	case <-b.closed:
	case b.userLeft <- userID:
	}
}

// EnterBot enters a bot to the battle.
func (b *Battle) EnterBot(bot *api.Bot) bool {
	select {
	case <-b.started:
	case <-b.closed:
		return false
	case b.botEntered <- bot:
	}
	return true
}

// LeaveBot leaves a bot from the battle.
func (b *Battle) LeaveBot(bot *api.Bot) bool {
	select {
	case <-b.started:
	case <-b.closed:
		return false
	case b.botLeft <- bot:
	}
	return true
}

func (b *Battle) run(ctx context.Context) {
	defer close(b.closed)

preLoop:
	for {
		select {
		case <-ctx.Done():
			return
		case <-b.startRequested:
			// TODO: Relevant users counting.
			if len(b.scene.Players()) > 0 {
				break preLoop
			}
		case userID := <-b.userEntered:
			b.enterUser(userID)
		case <-b.userLeft:
			break
		case bot := <-b.botEntered:
			b.enterBot(bot)
		case bot := <-b.botLeft:
			b.leaveBot(bot)
		}
	}
	b.start()
	close(b.started)

	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			if b.update(now) {
				break loop
			}
		case input := <-b.Gateway.input:
			b.onInputReceive(input)
		case userID := <-b.userEntered:
			b.reenterUser(userID)
		case userID := <-b.userLeft:
			b.leaveUser(userID)
		}
	}
	b.finish()
	close(b.finished)
}

func (b *Battle) start() {
	b.scene.Start(time.Now())
	b.Gateway.outputStart(b.scene.Players(), b.scene.StartedAt())
	for _, actor := range b.scene.Actors() {
		b.Gateway.outputActor(b.scene.UserPlayersByTeam(), actor)
	}
	b.scene.Event().AddActorAddEventListener(b.onActorAdd)
	b.scene.Event().AddActorMoveEventListener(b.onActorMove)
	b.scene.Event().AddActorChangeVisibilityEventListener(b.onActorChangeVisibility)
	b.scene.Event().AddActorRemoveEventListener(b.onActorRemove)
	b.scene.Event().AddActorUsePingerEventListener(b.onActorUsePinger)
	b.scene.Event().AddActorUpdateEquipmentEventListener(b.onActorUpdateEquipment)
}

func (b *Battle) update(now time.Time) bool {
	b.scene.Update(now)
	if debug.Debug {
		debug.Debugger.Update(b.scene.Stage(), debug.SortedSights(b.scene.SightsByTeam()))
	}
	return b.judge.isBattleFinished()
}

func (b *Battle) finish() {
	if winner := b.judge.winner(); winner != nil {
		b.Gateway.outputFinish(&winner.ID, b.scene.Now())
	} else {
		b.Gateway.outputFinish(nil, b.scene.Now())
	}
	if debug.Debug {
		debug.Debugger.Update(nil, nil)
	}
}

func (b *Battle) enterUser(userID int64) {
	if _, ok := b.scene.SubmarineByPlayerID(userID); !ok {
		index := len(b.scene.Players())
		startPos := b.getStartPosition(index)
		teamLayer := scene.GetTeamLayer(index + 1)
		user := scene.NewPlayer(userID, true, teamLayer, startPos)
		actor.NewSubmarine(b.scene, user)
	}
}

func (b *Battle) reenterUser(userID int64) {
	if s, ok := b.scene.SubmarineByPlayerID(userID); ok {
		players := scene.PlayerSlice{s.Player()}
		b.Gateway.outputStart(players, b.scene.StartedAt())
		for _, actor := range b.scene.Actors() {
			b.Gateway.outputActor(players.GroupByTeam(), actor)
		}
	}
}

func (b *Battle) leaveUser(userID int64) {
	if s, ok := b.scene.SubmarineByPlayerID(userID); ok {
		s.Event().EmitUserLeaveEvent()
	}
}

func (b *Battle) enterBot(bot *api.Bot) {
	index := len(b.scene.Players())
	startPos := b.getStartPosition(index)
	teamLayer := scene.GetTeamLayer(index + 1)
	player := scene.NewPlayer(bot.Id, false, teamLayer, startPos)
	player.AI = ai.NewSimpleAI(b.scene)
	actor.NewSubmarine(b.scene, player)
}

func (b *Battle) leaveBot(bot *api.Bot) {
	if s, ok := b.scene.SubmarineByPlayerID(bot.Id); ok {
		s.Destroy()
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
	s, ok := b.scene.SubmarineByPlayerID(input.userID)
	if !ok {
		return
	}
	switch m := input.message.(type) {
	case *battleAPI.AccelerationRequest:
		s.Event().EmitAccelerationRequestEvent(m)
	case *battleAPI.BrakeRequest:
		s.Event().EmitBrakeRequestEvent(m)
	case *battleAPI.TurnRequest:
		s.Event().EmitTurnRequestEvent(m)
	case *battleAPI.TorpedoRequest:
		s.Event().EmitTorpedoRequestEvent(m)
	case *battleAPI.PingerRequest:
		s.Event().EmitPingerRequestEvent(m)
	case *battleAPI.WatcherRequest:
		s.Event().EmitWatcherRequestEvent(m)
	}
}

func (b *Battle) onActorAdd(actor scene.Actor) {
	b.Gateway.outputActor(b.scene.UserPlayersByTeam(), actor)
}

func (b *Battle) onActorMove(actor scene.Actor) {
	b.Gateway.outputMovement(actor)
}

func (b *Battle) onActorChangeVisibility(actor scene.Actor, teamLayer navmesh.LayerMask) {
	b.Gateway.outputVisibility(b.scene.UserPlayersByTeam(), actor, teamLayer)
}

func (b *Battle) onActorRemove(actor scene.Actor) {
	b.Gateway.outputDestruction(actor)
}

func (b *Battle) onActorUsePinger(actor scene.Actor, finished bool) {
	b.Gateway.outputPinger(actor, finished)
}

func (b *Battle) onActorUpdateEquipment(actor scene.Actor, equipment *battleAPI.Equipment) {
	players := b.scene.UserPlayersByTeam()[actor.Player().TeamLayer]
	b.Gateway.outputEquipment(players, equipment)
}
