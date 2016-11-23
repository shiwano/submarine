package sight

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

func countLights(lm *LightMap) int {
	lightsCount := 0
	for _, lightsY := range lm.Lights {
		for _, light := range lightsY {
			if light != nil {
				lightsCount++
			}
		}
	}
	return lightsCount
}

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
				So(lm.MeshVersion, ShouldEqual, mesh.Version)
			})
		})

		Convey(".LoadLightMapFromJSONFile", func() {
			Convey("should load the light map from the specified JSON file", func() {
				lm, err := LoadLightMapFromJSONFile("../fixtures/lightmap.json")
				So(err, ShouldBeNil)
				So(countLights(lm), ShouldEqual, 313)
				So(lm.MeshVersion, ShouldEqual, mesh.Version)
			})
		})

		Convey(".LoadLightMapFromMessagePackFile", func() {
			Convey("should load the light map from the specified MessagePack file", func() {
				lm, err := LoadLightMapFromMessagePackFile("../fixtures/lightmap.mpac")
				So(err, ShouldBeNil)
				So(countLights(lm), ShouldEqual, 313)
				So(lm.MeshVersion, ShouldEqual, mesh.Version)
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

		Convey("#ToMessagePack", func() {
			Convey("should return MessagePack encoding of the light map", func() {
				lm := GenerateLightMap(navMesh, 1, 3)
				actualData, err := lm.ToMessagePack()
				expectedData, _ := ioutil.ReadFile("../fixtures/lightmap.mpac")

				So(err, ShouldBeNil)
				So(actualData, ShouldResemble, expectedData)
			})
		})

		Convey("#lightByNavMeshPoint", func() {
			lm := GenerateLightMap(navMesh, 1, 3)

			Convey("with a point of a light", func() {
				Convey("should light the specified cell point", func() {
					p := &vec2.T{0, 0}
					light := lm.lightByNavMeshPoint(p)
					So(light, ShouldNotBeNil)
				})
			})

			Convey("with a invalid point", func() {
				Convey("should light the specified cell point", func() {
					p := &vec2.T{1000, -1000}
					light := lm.lightByNavMeshPoint(p)
					So(light, ShouldBeNil)
				})
			})
		})
	})
}
