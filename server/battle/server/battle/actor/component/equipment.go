package component

import (
	"time"

	"github.com/shiwano/submarine/server/battle/server/battle/context"
)

// Equipment manages equipment items of the submarine.
type Equipment struct {
	torpedos []*equipmentItem
	pinger   *equipmentItem
}

// NewEquipment creates a Equipment.
func NewEquipment(params *context.SubmarineParams) *Equipment {
	e := new(Equipment)
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
