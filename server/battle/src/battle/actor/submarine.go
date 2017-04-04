package actor

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/actor/component"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
	"github.com/shiwano/submarine/server/battle/src/logger"

	"github.com/ungerik/go3d/float64/vec2"
)

type submarine struct {
	*actor
	isUsingPinger bool
	timer         *component.Timer
	equipment     *component.Equipment
}

// NewSubmarine creates a submarine.
func NewSubmarine(scn scene.Scene, user *scene.Player) scene.Actor {
	s := &submarine{
		actor: newActor(scn, user, user.SubmarineParams, user.StartPosition, 0),
	}
	s.timer = component.NewTimer(scn.Now())
	s.equipment = component.NewEquipment(s.ID(), user.SubmarineParams)

	s.event.AddCollideWithOtherActorEventListener(s.onCollideWithOtherActor)
	s.event.AddCollideWithStageEventListener(s.onCollideWithStage)
	s.event.AddAccelerationRequestEventListener(s.onAccelerationRequest)
	s.event.AddBrakeRequestEventListener(s.onBrakeRequest)
	s.event.AddTurnRequestEventListener(s.onTurnRequest)
	s.event.AddTorpedoRequestEventListener(s.onTorpedoRequest)
	s.event.AddPingerRequestEventListener(s.onPingerRequest)
	s.event.AddWatcherRequestEventListener(s.onWatcherRequest)
	s.event.AddUserLeaveEventListener(s.onUserLeave)

	s.scene.Event().AddActorUsePingerEventListener(s.onActorUsePinger)
	s.scene.Event().EmitActorCreateEvent(s)
	return s
}

func (s *submarine) Update() {
	s.timer.Update(s.scene.Now())

	if s.player.AI != nil {
		s.player.AI.Update(s)
	}
}

func (s *submarine) OnDestroy() {
	if s.isUsingPinger {
		s.finishToUsePinger()
	}
}

func (s *submarine) Submarine(layer navmesh.LayerMask) *battleAPI.ActorSubmarineObject {
	e := &battleAPI.ActorSubmarineObject{
		IsUsingPinger: s.isUsingPinger,
	}
	if layer == s.player.TeamLayer {
		e.Equipment = s.equipment.ToAPIType()
	}
	return e
}

func (s *submarine) onCollideWithOtherActor(actor scene.Actor, point vec2.T) {
	if actor.Player() != s.player {
		s.idle()
	}
}

func (s *submarine) onCollideWithStage(point vec2.T) {
	s.idle()
}

func (s *submarine) onAccelerationRequest(m *battleAPI.AccelerationRequest) {
	s.accelerate(m.Direction)
}

func (s *submarine) onBrakeRequest(m *battleAPI.BrakeRequest) {
	s.brake(m.Direction)
}

func (s *submarine) onTurnRequest(m *battleAPI.TurnRequest) {
	s.turn(m.Direction)
}

func (s *submarine) onTorpedoRequest(m *battleAPI.TorpedoRequest) {
	if s.equipment.TryConsumeTorpedo(s.scene.Now()) {
		logger.Log.Debugf("%v shoots a torpedo", s)
		s.scene.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		normalizedVelocity := s.motor.NormalizedVelocity()
		startOffsetValue := s.stageAgent.SizeRadius() * s.player.TorpedoParams.StartOffsetDistance
		startPoint := normalizedVelocity.Scale(startOffsetValue).Add(s.Position())
		newTorpedo(s.scene, s.player, startPoint, s.motor.Direction())
	}
}

func (s *submarine) onPingerRequest(m *battleAPI.PingerRequest) {
	if !s.isUsingPinger && s.equipment.Pinger.TryConsume(s.scene.Now()) {
		logger.Log.Debugf("%v uses pinger", s)
		s.scene.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		s.isUsingPinger = true
		s.scene.Event().EmitActorUsePingerEvent(s, false)
		s.timer.Register(s.player.SubmarineParams.PingerIntervalSeconds, s.finishToUsePinger)
	}
}

func (s *submarine) onWatcherRequest(m *battleAPI.WatcherRequest) {
	if s.equipment.Watcher.TryConsume(s.scene.Now()) {
		logger.Log.Debugf("%v uses watcher", s)
		s.scene.Event().EmitActorUpdateEquipmentEvent(s, s.equipment.ToAPIType())
		newWatcher(s.scene, s.player, s.Position(), s.motor.Direction())
	}
}

func (s *submarine) onUserLeave() {
	s.brake(s.motor.Direction())
}

func (s *submarine) onActorUsePinger(a scene.Actor, finished bool) {
	pingerTeam := a.Player().TeamLayer
	if s.Player().TeamLayer == pingerTeam {
		return
	}
	s.visibility.Set(pingerTeam, !finished)
}

func (s *submarine) finishToUsePinger() {
	if s.isUsingPinger {
		s.scene.Event().EmitActorUsePingerEvent(s, true)
		s.isUsingPinger = false
	}
}
