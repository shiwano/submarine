package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestAgent(t *testing.T) {
	Convey("Agent", t, func() {
		mesh, _ := LoadMeshFromJSONFile("fixtures/mesh.json")
		navmesh := New(mesh)
		agent := navmesh.CreateAgent(6, &vec2.Zero)

		Convey("should implement Object interface", func() {
			So(agent, ShouldImplement, (*Object)(nil))
		})

		Convey("#Move", func() {
			Convey("with the position which is in of the mesh", func() {
				Convey("should set the position", func() {
					agent.Move(&vec2.T{1, 2})
					So(agent.Position()[0], ShouldEqual, 1)
					So(agent.Position()[1], ShouldEqual, 2)
				})
			})

			Convey("with the position which is out of the mesh", func() {
				Convey("should not set the position", func() {
					agent.Move(&vec2.T{9999, 9999})
					So(agent.Position()[0], ShouldEqual, 0)
					So(agent.Position()[1], ShouldEqual, 0)

					agent.Move(&vec2.T{-8, 0})
					So(agent.Position()[0], ShouldEqual, 0)
					So(agent.Position()[1], ShouldEqual, 0)
				})
			})
		})
	})
}
