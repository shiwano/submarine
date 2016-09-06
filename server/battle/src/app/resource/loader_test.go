package resource

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoader(t *testing.T) {
	Convey("Loader", t, func() {
		loader := newLoader()

		Convey("#LoadStageMesh", func() {
			Convey("with valid stage code", func() {
				Convey("should return the specified stage mesh", func() {
					mesh, err := loader.LoadStageMesh(1)
					So(mesh, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})

			Convey("with invalid stage code", func() {
				Convey("should return nil", func() {
					mesh, err := loader.LoadStageMesh(-1)
					So(mesh, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
