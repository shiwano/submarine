package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLayerMask(t *testing.T) {
	Convey("LayerMask", t, func() {
		Convey("#Has", func() {
			l := Layer02 | Layer14 | Layer16

			Convey("with contained layer mask", func() {
				Convey("should return true", func() {
					So(l.Has(Layer14), ShouldBeTrue)
					So(l.Has(Layer02|Layer16), ShouldBeTrue)
				})
			})

			Convey("with no-contained layer mask", func() {
				Convey("should return false", func() {
					So(l.Has(Layer08), ShouldBeFalse)
				})
			})
		})

		Convey("#Set", func() {
			Convey("should set bit flag of the specified layer mask", func() {
				l := LayerMask(0)
				l.Set(Layer09)
				So(l.Has(Layer09), ShouldBeTrue)
			})
		})

		Convey("#Clear", func() {
			Convey("should clear bit flag of the specified layer mask", func() {
				l := Layer03 | Layer07
				l.Clear(Layer03)
				So(l.Has(Layer03), ShouldBeFalse)
			})
		})
	})
}
