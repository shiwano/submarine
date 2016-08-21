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
		agent := navmesh.CreateAgent(6, &vec2.T{1, 1})

		Convey("should implement Object interface", func() {
			So(agent, ShouldImplement, (*Object)(nil))
		})

		Convey("#Warp", func() {
			Convey("should set the position", func() {
				agent.Move(&vec2.T{2, 3}, 0)
				So(agent.Position()[0], ShouldEqual, 2)
				So(agent.Position()[1], ShouldEqual, 3)
			})
		})

		Convey("#Move", func() {
			Convey("with the position which is in the mesh", func() {
				Convey("should set the position", func() {
					agent.Move(&vec2.T{2, 3}, 0)
					So(agent.Position()[0], ShouldEqual, 2)
					So(agent.Position()[1], ShouldEqual, 3)
				})
			})

			Convey("with the position which is out of the mesh", func() {
				Convey("should set the intersection point", func() {
					agent.Move(&vec2.T{1, 9999}, 0)
					So(agent.Position()[0], ShouldEqual, 1)
					So(agent.Position()[1], ShouldEqual, 7)
				})

				Convey("should call the collide handler", func() {
					hitInfo := agent.Move(&vec2.T{1, 9999}, 0)
					So(hitInfo.Object, ShouldBeNil)
					So(hitInfo.Point[0], ShouldEqual, 1)
					So(hitInfo.Point[1], ShouldEqual, 7)
				})
			})

			Convey("with the position which collided with other object", func() {
				otherObj := navmesh.CreateAgent(2, &vec2.T{1, 6})
				otherObj.SetLayer(Layer02)

				Convey("should set the intersection point", func() {
					agent.Move(&vec2.T{1, 5}, 0)
					So(agent.Position()[0], ShouldEqual, 1)
					So(agent.Position()[1], ShouldEqual, 6-1)
				})

				Convey("should return the collided hitInfo", func() {
					hitInfo := agent.Move(&vec2.T{1, 5}, 0)
					So(hitInfo.Object, ShouldEqual, otherObj)
					So(hitInfo.Point[0], ShouldEqual, 1)
					So(hitInfo.Point[1], ShouldEqual, 6-1)
				})

				Convey("and with the other object's layer as ignoredLayer parameter", func() {
					Convey("should ignore the other object", func() {
						hitInfo := agent.Move(&vec2.T{1, 3}, Layer02)
						So(hitInfo, ShouldBeNil)
					})
				})
			})
		})
	})
}
