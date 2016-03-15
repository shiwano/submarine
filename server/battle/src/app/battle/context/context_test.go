package context

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContextTest(t *testing.T) {
	Convey("Context", t, func() {
		battleContext := NewContext()

		Convey("when an actor is created", func() {
			Convey("should add the actor", func() {
				actor := newSubmarine(battleContext)
				So(battleContext.HasActor(actor.ID()), ShouldBeTrue)
			})

			Convey("should call the actor's Start method", func() {
				actor := newSubmarine(battleContext)
				So(actor.isCalledStart, ShouldBeTrue)
			})

			Convey("should emit the ActorAdded event", func() {
				isCalled := false
				battleContext.Event.On(event.ActorAdd, func(a Actor) { isCalled = true })
				newSubmarine(battleContext)
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("when an actor is destroyed", func() {
			actor := newSubmarine(battleContext)

			Convey("should remove the actor", func() {
				actor.Destroy()
				So(battleContext.HasActor(actor.ID()), ShouldBeFalse)
			})

			Convey("should call the actor's OnDestroy method", func() {
				actor.Destroy()
				So(actor.isCalledOnDestroy, ShouldBeTrue)
			})

			Convey("should emit the ActorRemoved event", func() {
				isCalled := false
				battleContext.Event.On(event.ActorRemove, func(a Actor) { isCalled = true })
				actor.Destroy()
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("#Actor", func() {
			actorID := newSubmarine(battleContext).ID()

			Convey("with valid actor id", func() {
				Convey("should return the actor", func() {
					actor := battleContext.Actor(actorID)
					So(actor.ID(), ShouldEqual, actorID)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					actor := battleContext.Actor(actorID + 1)
					So(actor, ShouldBeNil)
				})
			})
		})

		Convey("#SubmarineByUserID", func() {
			userID := newSubmarine(battleContext).UserID()

			Convey("with valid user id", func() {
				Convey("should return the user's submarine", func() {
					submarine := battleContext.SubmarineByUserID(userID)
					So(submarine.UserID(), ShouldEqual, userID)
					So(submarine.ActorType(), ShouldEqual, battle.ActorType_Submarine)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					submarine := battleContext.SubmarineByUserID(userID + 1)
					So(submarine, ShouldBeNil)
				})
			})
		})
	})
}
