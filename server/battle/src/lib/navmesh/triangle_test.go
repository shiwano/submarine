package navmesh

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestTriangle(t *testing.T) {
	Convey("Triangle", t, func() {
		v1 := &vec2.T{5, 5}
		v2 := &vec2.T{-5, 5}
		v3 := &vec2.T{0, -5}
		triangle := newTriangle(v1, v2, v3)

		Convey("#containsPoint", func() {
			Convey("with a point in the triangle", func() {
				So(triangle.containsPoint(&vec2.T{0, 0}), ShouldBeTrue)
			})

			Convey("with a point out the triangle", func() {
				So(triangle.containsPoint(&vec2.T{0, 6}), ShouldBeFalse)
			})

			Convey("with a point on the triangle edge", func() {
				So(triangle.containsPoint(&vec2.T{-1, 5}), ShouldBeTrue)

				v1 := &vec2.T{5, 5}
				v2 := &vec2.T{0, 5}
				v3 := &vec2.T{0, -5}
				triangle := newTriangle(v1, v2, v3)
				So(triangle.containsPoint(&vec2.T{0, -1}), ShouldBeTrue)

				v1 = &vec2.T{0, 0}
				v2 = &vec2.T{-10, 7}
				v3 = &vec2.T{8, 7}
				triangle = newTriangle(v1, v2, v3)
				So(triangle.containsPoint(&vec2.T{-6, 7}), ShouldBeTrue)
			})

			Convey("with a point on the triangle vertex", func() {
				So(triangle.containsPoint(v1), ShouldBeTrue)
				So(triangle.containsPoint(v2), ShouldBeTrue)
				So(triangle.containsPoint(v3), ShouldBeTrue)
			})
		})

		Convey("#hasVertex", func() {
			Convey("with a vertex on the triangle", func() {
				So(triangle.hasVertex(triangle.Vertices[0]), ShouldBeTrue)
			})

			Convey("with a vertex not on the triangle", func() {
				So(triangle.hasVertex(&vec2.T{-1, -1}), ShouldBeFalse)
			})
		})

		Convey("#vertexIndex", func() {
			Convey("with a vertex on the triangle", func() {
				index, ok := triangle.vertexIndex(triangle.Vertices[0])
				So(ok, ShouldBeTrue)
				So(index, ShouldEqual, 0)
			})

			Convey("with a vertex not on the triangle", func() {
				index, ok := triangle.vertexIndex(&vec2.T{-1, -1})
				So(ok, ShouldBeFalse)
				So(index, ShouldEqual, -1)
			})
		})
	})
}
