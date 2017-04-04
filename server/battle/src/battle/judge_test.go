package battle

import (
	"testing"
	"time"

	"github.com/shiwano/submarine/server/battle/src/battle/actor"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"
	"github.com/shiwano/submarine/server/battle/src/resource"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestJudge(t *testing.T) {
	Convey("judge", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		lightMap, _ := resource.Loader.LoadLightMap(1)
		j := newJudge(scene.NewScene(stageMesh, lightMap), time.Second*10)

		Convey("#isBattleFinished", func() {
			now, _ := time.Parse(time.RFC3339, "2016-01-01T12:00:00+00:00")
			j.scene.Start(now)

			Convey("when the elapsed time is over the time limit", func() {
				Convey("should return the elapsed time since start of battle", func() {
					now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:10+00:00")
					j.scene.Update(now)

					So(j.isBattleFinished(), ShouldBeTrue)
				})
			})

			Convey("when the elapsed time is not over the time limit", func() {
				Convey("should return the elapsed time since start of battle", func() {
					now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:05+00:00")
					j.scene.Update(now)

					So(j.isBattleFinished(), ShouldBeFalse)
				})
			})
		})

		Convey("#winner", func() {
			user1 := scene.NewPlayer(1, true, scene.LayerTeam1, &vec2.T{0, 0})
			user2 := scene.NewPlayer(2, true, scene.LayerTeam1, &vec2.T{0, 0})
			actor.NewSubmarine(j.scene, user1)
			submarine2 := actor.NewSubmarine(j.scene, user2)

			Convey("when users count is 1", func() {
				Convey("should return the last user", func() {
					submarine2.Destroy()
					So(j.winner(), ShouldEqual, user1)
				})
			})

			Convey("when users count is over 1", func() {
				Convey("should return nil", func() {
					So(j.winner(), ShouldEqual, nil)
				})
			})
		})
	})
}
