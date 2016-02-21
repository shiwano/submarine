package battle

import (
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestActorContainerTest(t *testing.T) {
	Convey("ActorContainer", t, func() {
		context := newContext()
		container := newActorContainer(context)

		Convey("#createSubmarine", func() {
			Convey("should return the created user's submarine", func() {
				submarine := container.createSubmarine(1)
				So(submarine.UserID(), ShouldEqual, 1)
				So(submarine.ActorType(), ShouldEqual, battle.ActorType_Submarine)
			})

			Convey("should emit the actorCreated event", func() {
				var createdActor Actor
				context.event.On(event.ActorCreated, func(a Actor) { createdActor = a })
				submarine := container.createSubmarine(1)
				So(submarine.ID(), ShouldEqual, createdActor.ID())
			})
		})

		Convey("#getActor", func() {
			actorID := container.createSubmarine(1).ID()

			Convey("with valid actor id", func() {
				Convey("should return the actor", func() {
					actor := container.getActor(actorID)
					So(actor.ID(), ShouldEqual, actorID)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					actor := container.getActor(actorID + 1)
					So(actor, ShouldBeNil)
				})
			})
		})

		Convey("#getSubmarineByUserID", func() {
			userID := container.createSubmarine(1).UserID()

			Convey("with valid user id", func() {
				Convey("should return the user's submarine", func() {
					submarine := container.getSubmarineByUserID(userID)
					So(submarine.UserID(), ShouldEqual, userID)
					So(submarine.ActorType(), ShouldEqual, battle.ActorType_Submarine)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					submarine := container.getSubmarineByUserID(userID + 1)
					So(submarine, ShouldBeNil)
				})
			})
		})

		Convey("#destroyActor", func() {
			actor := container.createSubmarine(1)

			Convey("should destroy the actor", func() {
				container.destroyActor(actor)
				So(container.existsActor(actor.ID()), ShouldBeFalse)
			})

			Convey("should emit the actorDestroyed event", func() {
				var destroyedActor Actor
				context.event.On(event.ActorDestroyed, func(a Actor) { destroyedActor = a })
				container.destroyActor(actor)
				So(actor.ID(), ShouldEqual, destroyedActor.ID())
			})
		})
	})
}
