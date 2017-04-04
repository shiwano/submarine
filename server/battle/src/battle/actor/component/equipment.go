package component

import (
	"time"

	"github.com/shiwano/submarine/server/battle/lib/currentmillis"
	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
)

// Equipment manages equipment items of the submarine.
type Equipment struct {
	actorID  int64
	torpedos []*EquipmentItem
	Pinger   *EquipmentItem
	Watcher  *EquipmentItem
}

// NewEquipment creates a Equipment.
func NewEquipment(actorID int64, params *scene.SubmarineParams) *Equipment {
	e := new(Equipment)
	e.actorID = actorID
	for i := int64(0); i < params.TorpedoCount; i++ {
		e.torpedos = append(e.torpedos, &EquipmentItem{
			cooldownDuration: time.Duration(float64(time.Second) * params.TorpedoCooldownSeconds),
		})
	}
	e.Pinger = &EquipmentItem{
		cooldownDuration: time.Duration(float64(time.Second) * params.PingerCooldownSeconds),
	}
	e.Watcher = &EquipmentItem{
		cooldownDuration: time.Duration(float64(time.Second) * params.WatcherCooldownSeconds),
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
	message.Pinger = e.Pinger.toAPIType()
	message.Watcher = e.Watcher.toAPIType()
	return message
}

// TryConsumeTorpedo try to consume a torpedo item, and returns whether it succeed.
func (e *Equipment) TryConsumeTorpedo(now time.Time) bool {
	for _, t := range e.torpedos {
		if t.TryConsume(now) {
			return true
		}
	}
	return false
}

// EquipmentItem represents an equipment item.
type EquipmentItem struct {
	cooldownStartedAt time.Time
	cooldownDuration  time.Duration
}

func (i *EquipmentItem) toAPIType() *battleAPI.EquipmentItem {
	return &battleAPI.EquipmentItem{
		CooldownStartedAt: currentmillis.Millis(i.cooldownStartedAt),
		CooldownDuration:  currentmillis.DurationMillis(i.cooldownDuration),
	}
}

// TryConsume try to consume the item, and returns whether it succeed.
func (i *EquipmentItem) TryConsume(now time.Time) bool {
	if i.isConsumable(now) {
		i.cooldownStartedAt = now
		return true
	}
	return false
}

func (i *EquipmentItem) isConsumable(now time.Time) bool {
	return now.Sub(i.cooldownStartedAt) >= i.cooldownDuration
}
