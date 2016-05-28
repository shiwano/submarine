package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestEdge(t *testing.T) {
	Convey("Edge", t, func() {
		edge := &Edge{
			&vec2.T{5, 0},
			&vec2.T{-5, 0},
		}

		Convey("#intersect", func() {
			Convey("with an intersected points", func() {
				Convey("should return the intersection point", func() {
					result := edge.intersectWithLine(
						&vec2.T{0, 5},
						(&vec2.T{0, -5}).Sub(&vec2.T{0, 5}),
					)
					So(result[0], ShouldEqual, 0)
					So(result[1], ShouldEqual, 0)
				})
			})

			Convey("with an no-intersected points", func() {
				Convey("should return nil", func() {
					result := edge.intersectWithLine(
						&vec2.T{-1, 1},
						(&vec2.T{1, 1}).Sub(&vec2.T{-1, 1}),
					)
					So(result, ShouldBeNil)
				})
			})
		})
	})
}
