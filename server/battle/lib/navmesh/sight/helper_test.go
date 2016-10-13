package sight

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

func TestHelper(t *testing.T) {
	Convey("helper", t, func() {
		mesh, _ := navmesh.LoadMeshFromJSONFile("../fixtures/mesh.json")
		navMesh := navmesh.New(mesh)
		h := newHelper(navMesh, 1, 3)

		Convey(".newHelper", func() {
			Convey("should set NavMesh's information", func() {
				So(h.CellSize, ShouldEqual, 1)
				So(h.LightRange, ShouldEqual, 3)
				So(h.MinX, ShouldEqual, -10)
				So(h.MinY, ShouldEqual, -14)
				So(h.MaxX, ShouldEqual, 10)
				So(h.MaxY, ShouldEqual, 7)
				So(h.Width, ShouldEqual, 21)
				So(h.Height, ShouldEqual, 22)
			})
		})

		Convey("#convertNavMeshPointToCellPoint", func() {
			Convey("should return the relavant cell point", func() {
				p := h.cellPointByNavMeshPoint(&vec2.T{-10, -14})
				So(p, ShouldResemble, cellPoint{0, 0})

				p = h.cellPointByNavMeshPoint(&vec2.T{0, 0})
				So(p, ShouldResemble, cellPoint{10, 14})

				p = h.cellPointByNavMeshPoint(&vec2.T{10, 7})
				So(p, ShouldResemble, cellPoint{h.Width - 1, h.Height - 1})
			})
		})

		Convey("#convertCellPointToNavMeshPoint", func() {
			Convey("should return the relavant NavMesh point", func() {
				p := h.navMeshPointByCellPoint(&cellPoint{0, 0})
				So(p, ShouldResemble, &vec2.T{-10, -14})

				p = h.navMeshPointByCellPoint(&cellPoint{10, 14})
				So(p, ShouldResemble, &vec2.T{0, 0})

				p = h.navMeshPointByCellPoint(&cellPoint{h.Width - 1, h.Height - 1})
				So(p, ShouldResemble, &vec2.T{10, 7})
			})
		})
	})
}
