package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLayerMask(t *testing.T) {
	Convey("LayerMask", t, func() {
		Convey("#Has", func() {
			l := Layer1 | Layer13 | Layer15

			Convey("with contained layer mask", func() {
				Convey("should return true", func() {
					So(l.Has(Layer13), ShouldBeTrue)
					So(l.Has(Layer1|Layer15), ShouldBeTrue)
				})
			})

			Convey("with no-contained layer mask", func() {
				Convey("should return false", func() {
					So(l.Has(Layer7), ShouldBeFalse)
				})
			})
		})

		Convey("#Set", func() {
			Convey("should set bit flag of the specified layer mask", func() {
				l := LayerMask(0)
				l.Set(Layer8)
				So(l.Has(Layer8), ShouldBeTrue)
			})
		})

		Convey("#Clear", func() {
			Convey("should clear bit flag of the specified layer mask", func() {
				l := Layer2 | Layer6
				l.Clear(Layer2)
				So(l.Has(Layer2), ShouldBeFalse)
			})
		})
	})
}
