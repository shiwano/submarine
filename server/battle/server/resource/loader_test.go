package resource

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoader(t *testing.T) {
	Convey("Loader", t, func() {
		loader := newLoader()

		Convey("#LoadStageMesh", func() {
			Convey("with valid stage code", func() {
				Convey("should return the specified stage mesh", func() {
					mesh, err := loader.LoadMesh(1)
					So(mesh, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})

			Convey("with invalid stage code", func() {
				Convey("should return nil", func() {
					mesh, err := loader.LoadMesh(-1)
					So(mesh, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey("#LoadLightMap", func() {
			Convey("with valid stage code", func() {
				Convey("should return the specified light map", func() {
					lightMap, err := loader.LoadLightMap(1, 2, 3)
					So(lightMap, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
			})

			Convey("with invalid stage code", func() {
				Convey("should return nil", func() {
					lightMap, err := loader.LoadLightMap(-1, 2, 3)
					So(lightMap, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
