package battle

import (
	"testing"
	"time"

	"github.com/shiwano/submarine/server/battle/server/battle/actor"
	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/resource"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestJudge(t *testing.T) {
	Convey("judge", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		lightMap, _ := resource.Loader.LoadLightMap(1)
		j := newJudge(context.NewContext(stageMesh, lightMap), time.Second*10)

		Convey("#isBattleFinished", func() {
			j.context.StartedAt, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:00+00:00")

			Convey("when the elapsed time is over the time limit", func() {
				Convey("should return the elapsed time since start of battle", func() {
					j.context.Now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:10+00:00")
					So(j.isBattleFinished(), ShouldBeTrue)
				})
			})

			Convey("when the elapsed time is not over the time limit", func() {
				Convey("should return the elapsed time since start of battle", func() {
					j.context.Now, _ = time.Parse(time.RFC3339, "2016-01-01T12:00:05+00:00")
					So(j.isBattleFinished(), ShouldBeFalse)
				})
			})
		})

		Convey("#winner", func() {
			user1 := context.NewPlayer(1, true, context.LayerTeam1, &vec2.T{0, 0})
			user2 := context.NewPlayer(2, true, context.LayerTeam1, &vec2.T{0, 0})
			actor.NewSubmarine(j.context, user1)
			submarine2 := actor.NewSubmarine(j.context, user2)

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
