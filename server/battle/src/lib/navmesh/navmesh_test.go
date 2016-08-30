package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestNavMesh(t *testing.T) {
	Convey("mesh", t, func() {
		mesh, _ := LoadMeshFromJSONFile("fixtures/mesh.json")
		navmesh := New(mesh)

		Convey("#CreateAgent", func() {
			Convey("should creates an agent", func() {
				agent := navmesh.CreateAgent(1, &vec2.Zero)
				So(agent.ID(), ShouldEqual, 1)
				So(navmesh.Objects[1], ShouldEqual, agent)
			})
		})

		Convey("#destoryObject", func() {
			Convey("should destroys the specified object", func() {
				agent := navmesh.CreateAgent(1, &vec2.Zero)
				navmesh.destroyObject(agent.ID())
				So(navmesh.Objects, ShouldNotContainKey, (int64)(1))
			})
		})

		Convey("#ContainsPoint", func() {
			Convey("with a contained point", func() {
				Convey("should return true", func() {
					So(navmesh.ContainsPoint(&vec2.T{0, 0}), ShouldBeTrue)
				})
			})

			Convey("with a no-contained point", func() {
				Convey("should return false", func() {
					So(navmesh.ContainsPoint(&vec2.T{999, 999}), ShouldBeFalse)
				})
			})
		})

		Convey("#Raycast", func() {
			Convey("with ray parameters which intersected with the mesh", func() {
				Convey("should return the intersection point", func() {
					hitInfo := navmesh.Raycast(
						&vec2.T{1, 0},
						(&vec2.T{1, 100}).Sub(&vec2.T{1, 0}),
						0,
					)
					So(hitInfo.Object, ShouldBeNil)
					So(hitInfo.Point[0], ShouldEqual, 1)
					So(hitInfo.Point[1], ShouldEqual, 7)
				})
			})

			Convey("with ray parameters which intersected with an object", func() {
				Convey("should return the intersection object and point", func() {
					agent := navmesh.CreateAgent(2, &vec2.T{1, 3})
					hitInfo := navmesh.Raycast(
						&vec2.T{1, 0},
						(&vec2.T{1, 100}).Sub(&vec2.T{1, 0}),
						0,
					)
					So(hitInfo.Object.ID(), ShouldEqual, agent.ID())
					So(hitInfo.Point[0], ShouldEqual, 1)
					So(hitInfo.Point[1], ShouldEqual, 2)
				})
			})

			Convey("with ray parameters which did not intersect", func() {
				Convey("should return nil", func() {
					hitInfo := navmesh.Raycast(
						&vec2.T{1, 100},
						(&vec2.T{1, 200}).Sub(&vec2.T{1, 100}),
						0,
					)
					So(hitInfo, ShouldBeNil)
				})
			})

			Convey("with an ignoredLayer", func() {
				Convey("should ignore objects that has the specified layer", func() {
					agent := navmesh.CreateAgent(2, &vec2.T{1, 3})
					agent.SetLayer(Layer02)
					hitInfo := navmesh.Raycast(
						&vec2.T{1, 0},
						(&vec2.T{1, 100}).Sub(&vec2.T{1, 0}),
						Layer02,
					)
					So(hitInfo.Object, ShouldBeNil)
				})
			})
		})

		Convey("#FindPath", func() {
			Convey("with points that are inside the same triangle", func() {
				Convey("should return the path", func() {
					start := &vec2.T{-7, 3.5}
					goal := &vec2.T{-6, 2}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{
						vec2.T{-7, 3.5},
						vec2.T{-6, 2},
					})
				})
			})

			Convey("with points that are outside the mesh", func() {
				Convey("should return the empty path", func() {
					start := &vec2.T{9999, 9999}
					goal := &vec2.T{-9999, -9999}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{})

					start = &vec2.T{-7, 3.5}
					goal = &vec2.T{-9999, -9999}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{})

					start = &vec2.T{9999, 9999}
					goal = &vec2.T{-6, 2}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{})
				})
			})

			Convey("with points that can make a zigzag path", func() {
				Convey("should return the path", func() {
					start := &vec2.T{-7, 3.5}
					goal := &vec2.T{1, -11}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{
						vec2.T{-7, 3.5},
						vec2.T{0, 0},
						vec2.T{0, -9},
						vec2.T{1, -11},
					})
				})
			})

			Convey("with points that can make a straight path", func() {
				Convey("should return the path", func() {
					start := &vec2.T{3, 2}
					goal := &vec2.T{2, -10}
					So(navmesh.FindPath(start, goal), ShouldResemble, []vec2.T{
						vec2.T{3, 2},
						vec2.T{2, -10},
					})
				})
			})
		})
	})
}
