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
	helper := newHelper(navMesh, 1, 3)

	Convey("Light", t, func() {
		Convey("newLight", func() {
			Convey("should return the light that has relevant lit points", func() {
				l := newLight(navMesh, helper, &vec2.T{0, 1})
				So(l.isLighting(), ShouldBeTrue)
				So(l.LitPoints, ShouldHaveLength, 25)
			})
		})
	})
}
