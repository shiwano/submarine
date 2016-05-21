package resource

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoader(t *testing.T) {
	Convey("Loader", t, func() {
		loader := newLoader()

		Convey("#LoadMap", func() {
			Convey("with valid map code", func() {
				Convey("should return the specified battle map", func() {
					battleMap, err := loader.LoadBattleMap(1)
					So(battleMap.Code, ShouldEqual, 1)
					So(battleMap.NavMesh.Mesh.Triangles, ShouldHaveLength, 6)
					So(err, ShouldBeNil)
				})
			})

			Convey("with invalid map code", func() {
				Convey("should return nil", func() {
					battleMap, err := loader.LoadBattleMap(-1)
					So(battleMap, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
