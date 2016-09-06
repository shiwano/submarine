package sight

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"testing"
)

func TestHelper(t *testing.T) {
	Convey("helper", t, func() {
		mesh, _ := navmesh.LoadMeshFromJSONFile("../fixtures/mesh.json")
		navMesh := navmesh.New(mesh)
		h := newHelper(navMesh, 1, 3)

		Convey(".newHelper", func() {
			Convey("should set NavMesh's information", func() {
				So(h.cellSize, ShouldEqual, 1)
				So(h.lightRange, ShouldEqual, 3)
				So(h.lightRangeSqr, ShouldEqual, 9)
				So(h.lightDiameter, ShouldEqual, 7)
				So(h.minX, ShouldEqual, -10)
				So(h.minY, ShouldEqual, -14)
				So(h.maxX, ShouldEqual, 10)
				So(h.maxY, ShouldEqual, 7)
				So(h.width, ShouldEqual, 21)
				So(h.height, ShouldEqual, 22)
			})
		})

		Convey("#convertNavMeshPointToCellPoint", func() {
			Convey("should return the relavant cell point", func() {
				p := h.cellPointByNavMeshPoint(&vec2.T{-10, -14})
				So(p, ShouldResemble, cellPoint{0, 0})

				p = h.cellPointByNavMeshPoint(&vec2.T{0, 0})
				So(p, ShouldResemble, cellPoint{10, 14})

				p = h.cellPointByNavMeshPoint(&vec2.T{10, 7})
				So(p, ShouldResemble, cellPoint{h.width - 1, h.height - 1})
			})
		})

		Convey("#convertCellPointToNavMeshPoint", func() {
			Convey("should return the relavant NavMesh point", func() {
				p := h.navMeshPointByCellPoint(&cellPoint{0, 0})
				So(p, ShouldResemble, &vec2.T{-10, -14})

				p = h.navMeshPointByCellPoint(&cellPoint{10, 14})
				So(p, ShouldResemble, &vec2.T{0, 0})

				p = h.navMeshPointByCellPoint(&cellPoint{h.width - 1, h.height - 1})
				So(p, ShouldResemble, &vec2.T{10, 7})
			})
		})
	})
}
