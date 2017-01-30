package component

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/server/battle/context"
)

func TestEquipment(t *testing.T) {
	Convey("Equipment", t, func() {
		startTime := time.Now()
		params := &context.SubmarineParams{
			TorpedoCount:           2,
			TorpedoCooldownSeconds: 10,
			PingerCooldownSeconds:  20,
			WatcherCooldownSeconds: 20,
		}
		e := NewEquipment(1, params)

		Convey("#ToApiType", func() {
			Convey("should returns a Equipment message", func() {
				message := e.ToAPIType()
				So(message.ActorId, ShouldEqual, 1)
				So(message.Coerce(), ShouldBeNil)
			})
		})

		Convey("#TryConsumeTorpedo", func() {
			Convey("should consume torpedo", func() {
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeFalse)
			})

			Convey("should start cooldown of the consumed torpedo", func() {
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime), ShouldBeTrue)
				So(e.TryConsumeTorpedo(startTime.Add(time.Second*10)), ShouldBeTrue)
			})
		})
	})
}

func TestEquipmentItem(t *testing.T) {
	Convey("EquipmentItem", t, func() {
		startTime := time.Now()
		i := &EquipmentItem{cooldownDuration: time.Second * 20}

		Convey("#TryConsume", func() {
			Convey("should consume the item", func() {
				So(i.TryConsume(startTime), ShouldBeTrue)
				So(i.TryConsume(startTime), ShouldBeFalse)
			})

			Convey("should start cooldown of the consumed item", func() {
				So(i.TryConsume(startTime), ShouldBeTrue)
				So(i.TryConsume(startTime.Add(time.Second*10)), ShouldBeFalse)
				So(i.TryConsume(startTime.Add(time.Second*20)), ShouldBeTrue)
			})
		})
	})
}
