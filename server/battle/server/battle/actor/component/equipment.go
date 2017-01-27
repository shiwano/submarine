package component

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
)

// Equipment manages equipment items of the submarine.
type Equipment struct {
	actorID  int64
	torpedos []*equipmentItem
	pinger   *equipmentItem
}

// NewEquipment creates a Equipment.
func NewEquipment(actorID int64, params *context.SubmarineParams) *Equipment {
	e := new(Equipment)
	e.actorID = actorID
	for i := int64(0); i < params.TorpedoCount; i++ {
		e.torpedos = append(e.torpedos, &equipmentItem{
			cooldownDuration: time.Duration(float64(time.Second) * params.TorpedoCooldownSeconds),
		})
	}
	e.pinger = &equipmentItem{
		cooldownDuration: time.Duration(float64(time.Second) * params.PingerCooldownSeconds),
	}
	return e
}

// ToAPIType returns a Equipment message.
func (e *Equipment) ToAPIType() *battleAPI.Equipment {
	message := new(battleAPI.Equipment)
	message.ActorId = e.actorID
	for _, i := range e.torpedos {
		message.Torpedos = append(message.Torpedos, i.toAPIType())
	}
	message.Pinger = e.pinger.toAPIType()
	return message
}

// TryConsumeTorpedo try to consume a torpedo item, and returns whether it succeed.
func (e *Equipment) TryConsumeTorpedo(now time.Time) bool {
	for _, t := range e.torpedos {
		if t.tryConsume(now) {
			return true
		}
	}
	return false
}

// TryConsumePinger try to consume a pinger item, and returns whether it succeed.
func (e *Equipment) TryConsumePinger(now time.Time) bool {
	return e.pinger.tryConsume(now)
}

type equipmentItem struct {
	cooldownStartedAt time.Time
	cooldownDuration  time.Duration
}

func (i *equipmentItem) toAPIType() *battleAPI.EquipmentItem {
	return &battleAPI.EquipmentItem{
		CooldownStartedAt: currentmillis.Millis(i.cooldownStartedAt),
		CooldownDuration:  currentmillis.DurationMillis(i.cooldownDuration),
	}
}

func (i *equipmentItem) tryConsume(now time.Time) bool {
	if i.isConsumable(now) {
		i.cooldownStartedAt = now
		return true
	}
	return false
}

func (i *equipmentItem) isConsumable(now time.Time) bool {
	return now.Sub(i.cooldownStartedAt) >= i.cooldownDuration
}
