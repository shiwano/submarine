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

		Convey("#findPath", func() {
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
