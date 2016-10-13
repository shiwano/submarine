package ai

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestNavigator(t *testing.T) {
	Convey("navigator", t, func() {
		n := &navigator{}

		Convey("#start", func() {
			Convey("with a valid path", func() {
				Convey("should start navigation", func() {
					path := []vec2.T{vec2.T{0, 0}, vec2.T{1, 1}}
					n.start(path, &vec2.T{0, 0})
					So(n.isStarted(), ShouldBeTrue)
				})
			})

			Convey("with a invalid path that does not have enough length", func() {
				Convey("should stop navigation", func() {
					path := []vec2.T{vec2.T{0, 0}}
					n.start(path, &vec2.T{0, 0})
					So(n.isStarted(), ShouldBeFalse)
				})
			})
		})

		Convey("#stop", func() {
			path := []vec2.T{vec2.T{0, 0}, vec2.T{1, 1}}
			n.start(path, &vec2.T{0, 0})

			Convey("should stop navigation", func() {
				n.stop()
				So(n.isStarted(), ShouldBeFalse)
			})
		})

		Convey("#navigate", func() {
			Convey("when navigation stopped", func() {
				Convey("should return false", func() {
					ok, _ := n.navigate(&vec2.T{0, 0})
					So(ok, ShouldBeFalse)
				})
			})

			Convey("when navigation started", func() {
				path := []vec2.T{vec2.T{0, 0}, vec2.T{0, 10}, vec2.T{10, 10}}
				n.start(path, &vec2.T{0, 0})

				Convey("with a point that does not go through the next point", func() {
					Convey("should return the next point direction", func() {
						ok, direction := n.navigate(&vec2.T{0, 0})
						So(ok, ShouldBeTrue)
						So(math.Floor(direction), ShouldEqual, 90)

						ok, direction = n.navigate(&vec2.T{0, 1})
						So(ok, ShouldBeTrue)
						So(math.Floor(direction), ShouldEqual, 90)
					})
				})

				Convey("with a point that went through the next point", func() {
					ok, direction := n.navigate(&vec2.T{0, 11})

					Convey("should return the next point direction", func() {
						So(ok, ShouldBeTrue)
						So(math.Floor(direction), ShouldEqual, 354)
					})
					Convey("should change the next point", func() {
						So(n.nextPointIndex, ShouldEqual, 2)
					})
				})

				Convey("with a point that finished path navigation", func() {
					n.navigate(&vec2.T{0, 10})
					ok, _ := n.navigate(&vec2.T{10, 10})

					Convey("should return false", func() {
						So(ok, ShouldBeFalse)
					})
					Convey("should stop navigation", func() {
						So(n.isStarted(), ShouldBeFalse)
					})
				})
			})
		})
	})
}
