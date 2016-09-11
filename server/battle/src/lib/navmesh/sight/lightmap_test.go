package sight

import (
	"io/ioutil"
	"lib/navmesh"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLightMap(t *testing.T) {
	Convey("LightMap", t, func() {
		mesh, _ := navmesh.LoadMeshFromJSONFile("../fixtures/mesh.json")
		navMesh := navmesh.New(mesh)

		Convey(".GenerateLightMap", func() {
			Convey("should return a light map that has pre-calculated lights", func() {
				lm := GenerateLightMap(navMesh, 1, 3)
				lightsCount := 0
				for _, lightsY := range lm.Lights {
					for _, light := range lightsY {
						if light != nil {
							lightsCount++
						}
					}
				}
				So(lightsCount, ShouldEqual, 313)
			})
		})

		Convey(".LoadLightMapFromJSONFile", func() {
			Convey("should load the light map from the specified JSON file", func() {
				lm, err := LoadLightMapFromJSONFile("../fixtures/lightmap.json")
				So(err, ShouldBeNil)

				lightsCount := 0
				for _, lightsY := range lm.Lights {
					for _, light := range lightsY {
						if light != nil {
							lightsCount++
						}
					}
				}
				So(lightsCount, ShouldEqual, 313)
			})
		})

		Convey("#ToJSON", func() {
			Convey("should return JSON encoding of the light map", func() {
				lm := GenerateLightMap(navMesh, 1, 3)
				actualData, err := lm.ToJSON()
				expectedData, _ := ioutil.ReadFile("../fixtures/lightmap.json")

				So(err, ShouldBeNil)
				So(actualData, ShouldResemble, expectedData)
			})
		})
	})
}
