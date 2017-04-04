package scene

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
	"github.com/shiwano/submarine/server/battle/src/resource"
)

func TestSceneTest(t *testing.T) {
	Convey("scene", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		lightMap, _ := resource.Loader.LoadLightMap(1)
		scn := NewScene(stageMesh, lightMap)

		Convey("when an actor is created", func() {
			Convey("should add the actor", func() {
				actor := newSubmarine(scn, true)
				So(scn.HasActor(actor.ID()), ShouldBeTrue)
			})

			Convey("should call the actor's Start method", func() {
				actor := newSubmarine(scn, true)
				So(actor.isCalledStart, ShouldBeTrue)
			})

			Convey("should emit the ActorAdded event", func() {
				isCalled := false
				scn.Event().AddActorAddEventListener(func(a Actor) { isCalled = true })
				newSubmarine(scn, true)
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("when an actor is destroyed", func() {
			actor := newSubmarine(scn, true)
			newSubmarine(scn, false)
			newSubmarine(scn, true)

			Convey("should remove the actor", func() {
				actor.Destroy()
				So(scn.HasActor(actor.ID()), ShouldBeFalse)
				So(scn.Actors(), ShouldHaveLength, 2)
				So(scn.Players(), ShouldHaveLength, 2)
				So(scn.UserPlayersByTeam(), ShouldHaveLength, 1)
				So(scn.UserPlayersByTeam(), ShouldHaveLength, 1)
			})

			Convey("should call the actor's OnDestroy method", func() {
				actor.Destroy()
				So(actor.isCalledOnDestroy, ShouldBeTrue)
			})

			Convey("should emit the ActorRemoved event", func() {
				isCalled := false
				scn.Event().AddActorRemoveEventListener(func(a Actor) { isCalled = true })
				actor.Destroy()
				So(isCalled, ShouldBeTrue)
			})
		})

		Convey("#ElapsedTime", func() {
			now, _ := time.Parse(time.RFC3339, "2016-01-01T12:00:00+00:00")
			scn.Start(now)
			now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:40+00:00")
			scn.Update(now)

			Convey("should return the elapsed time since start of battle", func() {
				So(scn.ElapsedTime(), ShouldEqual, time.Second*40)
			})
		})

		Convey("#Actor", func() {
			actorID := newSubmarine(scn, true).ID()

			Convey("with valid actor id", func() {
				Convey("should return the actor", func() {
					a, ok := scn.Actor(actorID)
					So(ok, ShouldBeTrue)
					So(a.ID(), ShouldEqual, actorID)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					a, ok := scn.Actor(actorID + 1)
					So(ok, ShouldBeFalse)
					So(a, ShouldBeNil)
				})
			})
		})

		Convey("#SubmarineByUserID", func() {
			userID := newSubmarine(scn, true).Player().ID

			Convey("with valid user id", func() {
				Convey("should return the user's submarine", func() {
					s, ok := scn.SubmarineByPlayerID(userID)
					So(ok, ShouldBeTrue)
					So(s.Player().ID, ShouldEqual, userID)
					So(s.Type(), ShouldEqual, battleAPI.ActorType_Submarine)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					s, ok := scn.SubmarineByPlayerID(userID + 1)
					So(ok, ShouldBeFalse)
					So(s, ShouldBeNil)
				})
			})
		})
	})
}
