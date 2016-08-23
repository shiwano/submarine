package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestEdge(t *testing.T) {
	Convey("Edge", t, func() {
		e := newEdge(
			&vec2.T{5, 0},
			&vec2.T{-5, 0},
		)

		Convey("#containsPoint", func() {
			Convey("with an contained point", func() {
				Convey("should return true", func() {
					So(e.containsPoint(&vec2.T{0, 0}), ShouldBeTrue)
				})
			})

			Convey("with an no-contained point", func() {
				Convey("should return false", func() {
					So(e.containsPoint(&vec2.T{0, 1}), ShouldBeFalse)
				})
			})
		})

		Convey("#intersectWithLineSeg", func() {
			Convey("with an intersected points", func() {
				Convey("should return the intersection point", func() {
					p1 := &vec2.T{0, 5}
					p2 := &vec2.T{0, -5}

					p, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, 0)
					So(p[1], ShouldEqual, 0)
				})
			})

			Convey("with an no-intersected points", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{-1, 1}
					p2 := &vec2.T{1, 1}

					_, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
