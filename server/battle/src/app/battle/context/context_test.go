package context

import (
	"app/battle/event"
	"app/resource"
	battleAPI "app/typhenapi/type/submarine/battle"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContextTest(t *testing.T) {
	Convey("Context", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		c := NewContext(stageMesh)

		Convey("when an actor is created", func() {
			Convey("should add the actor", func() {
				actor := newSubmarine(c)
				So(c.HasActor(actor.ID()), ShouldBeTrue)
			})

			Convey("should call the actor's Start method", func() {
				actor := newSubmarine(c)
				So(actor.isCalledStart, ShouldBeTrue)
			})

			Convey("should emit the ActorAdded event", func() {
				isCalled := false
				c.Event.On(event.ActorAdd, func(a Actor) { isCalled = true })
				newSubmarine(c)
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("when an actor is destroyed", func() {
			actor := newSubmarine(c)
			newSubmarine(c)

			Convey("should remove the actor", func() {
				actor.Destroy()
				So(c.HasActor(actor.ID()), ShouldBeFalse)
				So(c.Actors(), ShouldHaveLength, 1)
				So(c.Players(), ShouldHaveLength, 1)
			})

			Convey("should call the actor's OnDestroy method", func() {
				actor.Destroy()
				So(actor.isCalledOnDestroy, ShouldBeTrue)
			})

			Convey("should emit the ActorRemoved event", func() {
				isCalled := false
				c.Event.On(event.ActorRemove, func(a Actor) { isCalled = true })
				actor.Destroy()
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("#ElapsedTime", func() {
			c.StartedAt, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:00+00:00")
			c.Now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:40+00:00")

			Convey("should return the elapsed time since start of battle", func() {
				So(c.ElapsedTime(), ShouldEqual, time.Second*40)
			})
		})

		Convey("#Actor", func() {
			actorID := newSubmarine(c).ID()

			Convey("with valid actor id", func() {
				Convey("should return the actor", func() {
					actor := c.Actor(actorID)
					So(actor.ID(), ShouldEqual, actorID)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					actor := c.Actor(actorID + 1)
					So(actor, ShouldBeNil)
				})
			})
		})

		Convey("#SubmarineByUserID", func() {
			userID := newSubmarine(c).Player().ID

			Convey("with valid user id", func() {
				Convey("should return the user's submarine", func() {
					submarine := c.SubmarineByPlayerID(userID)
					So(submarine.Player().ID, ShouldEqual, userID)
					So(submarine.Type(), ShouldEqual, battleAPI.ActorType_Submarine)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					submarine := c.SubmarineByPlayerID(userID + 1)
					So(submarine, ShouldBeNil)
				})
			})
		})
	})
}
