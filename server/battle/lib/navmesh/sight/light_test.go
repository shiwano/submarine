package sight

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

func TestLight(t *testing.T) {
	mesh, _ := navmesh.LoadMeshFromJSONFile("../fixtures/mesh.json")
	navMesh := navmesh.New(mesh)
	helper := newHelper(navMesh, 1, 2)

	Convey("Light", t, func() {
		Convey("newLight", func() {
			Convey("should return the light that has relevant lit points", func() {
				l := newLight(navMesh, helper, &vec2.T{0, 0})
				So(l.isLighting(), ShouldBeTrue)
				So(l.LitPoints, ShouldResemble, []cellPoint{
					cellPoint{8, 14},
					cellPoint{9, 14},
					cellPoint{9, 15},
					cellPoint{10, 12},
					cellPoint{10, 13},
					cellPoint{10, 14},
					cellPoint{10, 15},
					cellPoint{10, 16},
					cellPoint{11, 13},
					cellPoint{11, 14},
					cellPoint{11, 15},
					cellPoint{12, 14},
				})
			})
		})
	})
}
