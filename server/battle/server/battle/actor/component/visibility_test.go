package component

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVisibility(t *testing.T) {
	Convey("Visibility", t, func() {
		v := new(Visibility)

		Convey("#Set", func() {
			Convey("with true", func() {
				Convey("should set visibility", func() {
					v.Set(true)
					So(v.IsVisible(), ShouldBeTrue)
				})

				Convey("should set visibility in duplicate", func() {
					v.Set(true)
					v.Set(true)
					v.Set(false)
					So(v.IsVisible(), ShouldBeTrue)
				})
			})

			Convey("with false", func() {
				Convey("should unset visibility", func() {
					v.Set(true)
					v.Set(false)
					So(v.IsVisible(), ShouldBeFalse)
				})

				Convey("should not unset visibility in duplicate", func() {
					v.Set(true)
					v.Set(false)
					v.Set(false)
					v.Set(true)
					So(v.IsVisible(), ShouldBeTrue)
				})
			})

			Convey("when visibility changed", func() {
				Convey("should call ChangeHandler", func() {
					callCount := 0
					v.ChangeHandler = func() {
						callCount++
					}
					v.Set(true)
					v.Set(true)
					v.Set(false)
					v.Set(false)
					v.Set(false)
					So(callCount, ShouldEqual, 2)
				})
			})
		})
	})
}
