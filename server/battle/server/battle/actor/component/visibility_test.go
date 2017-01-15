package component

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

func TestVisibility(t *testing.T) {
	Convey("Visibility", t, func() {
		v := NewVisibility()

		Convey("#Set", func() {
			Convey("with true", func() {
				Convey("should set visibility of the specified layer", func() {
					v.Set(navmesh.Layer01, true)
					So(v.IsVisibleFrom(navmesh.Layer01), ShouldBeTrue)
					So(v.IsVisibleFrom(navmesh.Layer02), ShouldBeFalse)
				})

				Convey("should set visibility of the specified layer in duplicate", func() {
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, false)
					So(v.IsVisibleFrom(navmesh.Layer01), ShouldBeTrue)
				})
			})

			Convey("with false", func() {
				Convey("should unset visibility of the specified layer", func() {
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, false)
					So(v.IsVisibleFrom(navmesh.Layer01), ShouldBeFalse)
				})

				Convey("should not unset visibility of the specified layer in duplicate", func() {
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, false)
					v.Set(navmesh.Layer01, false)
					v.Set(navmesh.Layer01, true)
					So(v.IsVisibleFrom(navmesh.Layer01), ShouldBeTrue)
				})
			})

			Convey("when visibility changed", func() {
				Convey("should call ChangeHandler", func() {
					callCount := 0
					v.ChangeHandler = func(layer navmesh.LayerMask) {
						if layer == navmesh.Layer01 {
							callCount++
						}
					}
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, true)
					v.Set(navmesh.Layer01, false)
					v.Set(navmesh.Layer01, false)
					v.Set(navmesh.Layer01, false)
					So(callCount, ShouldEqual, 2)
				})
			})
		})
	})
}
