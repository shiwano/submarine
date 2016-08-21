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
					called := false
					agent.SetCollideHandler(func(obj Object, point vec2.T) {
						So(obj, ShouldBeNil)
						So(point[0], ShouldEqual, 1)
						So(point[1], ShouldEqual, 7)
						called = true
					})
					agent.Move(&vec2.T{1, 9999}, 0)
					So(called, ShouldBeTrue)
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

				Convey("should call the collide handler", func() {
					called := false
					agent.SetCollideHandler(func(obj Object, point vec2.T) {
						So(obj, ShouldEqual, otherObj)
						So(point[0], ShouldEqual, 1)
						So(point[1], ShouldEqual, 6-1)
						called = true
					})
					agent.Move(&vec2.T{1, 5}, 0)
					So(called, ShouldBeTrue)
				})

				Convey("should call the other object's collide handler", func() {
					called := false
					otherObj.SetCollideHandler(func(obj Object, point vec2.T) {
						So(obj, ShouldEqual, agent)
						So(point[0], ShouldEqual, 1)
						So(point[1], ShouldEqual, 6-1)
						called = true
					})
					agent.Move(&vec2.T{1, 5}, 0)
					So(called, ShouldBeTrue)
				})

				Convey("with the other object's layer as ignoredLayer parameter", func() {
					Convey("should ignore the other object", func() {
						called := false
						otherObj.SetCollideHandler(func(obj Object, point vec2.T) {
							called = true
						})
						agent.Move(&vec2.T{1, 3}, Layer02)
						So(called, ShouldBeFalse)
					})
				})
			})
		})
	})
}
