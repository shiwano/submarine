package navmesh

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestEdge(t *testing.T) {
	Convey("edge", t, func() {
		v1 := &vec2.T{5, 0}
		v2 := &vec2.T{-5, 0}
		v3 := &vec2.T{0, 5}
		triangle := newTriangle(v1, v2, v3)

		e := newEdge(triangle, 0, 1)

		Convey("#isEndPoint", func() {
			Convey("with a same point with the end point", func() {
				Convey("should return true", func() {
					So(e.isEndPoint(&vec2.T{5, 0}), ShouldBeTrue)
				})
			})

			Convey("with a different point with the end point", func() {
				Convey("should return false", func() {
					So(e.isEndPoint(&vec2.T{0, 0}), ShouldBeFalse)
				})
			})
		})

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
			Convey("with an intersected line segment", func() {
				Convey("should return the intersection point", func() {
					p1 := &vec2.T{0, 5}
					p2 := &vec2.T{0, -5}

					p, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, 0)
					So(p[1], ShouldEqual, 0)
				})
			})

			Convey("with an no-intersected line segment", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{-1, 1}
					p2 := &vec2.T{1, 1}

					_, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)
				})
			})

			Convey("with an parallel line segment", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{4, 0}
					p2 := &vec2.T{-4, 0}

					_, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)
				})
			})

			Convey("with an line segment whose line origin is on the edge", func() {
				Convey("and line vector faces the inside of the edge", func() {
					Convey("should return nil", func() {
						p1 := &vec2.T{0, 0}
						p2 := &vec2.T{0, 5}

						_, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
						So(ok, ShouldBeFalse)
					})
				})

				Convey("and line vector faces the outside of the edge", func() {
					Convey("should return the intersection point", func() {
						p1 := &vec2.T{0, 0}
						p2 := &vec2.T{0, -5}

						p, ok := e.intersectWithLineSeg(p1, p2.Sub(p1))
						So(ok, ShouldBeTrue)
						So(p[0], ShouldEqual, 0)
						So(p[1], ShouldEqual, 0)
					})
				})
			})
		})
	})
}
