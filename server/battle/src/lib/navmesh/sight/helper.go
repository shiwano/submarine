package sight

import (
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"math"
)

type cellPoint [2]int

type helper struct {
	cellSize float64

	lightRange    float64
	lightRangeSqr float64
	lightDiameter float64

	minX int
	minY int
	maxX int
	maxY int

	width  int
	height int
}

func newHelper(navMesh *navmesh.NavMesh, cellSize float64, lightRange float64) *helper {
	h := &helper{
		cellSize:      cellSize,
		lightRange:    lightRange,
		lightRangeSqr: lightRange * lightRange,
		lightDiameter: lightRange*2 + 1,
		minX:          int(math.Floor(navMesh.Mesh.Rect.Min[0] / cellSize)),
		minY:          int(math.Floor(navMesh.Mesh.Rect.Min[1] / cellSize)),
		maxX:          int(math.Ceil(navMesh.Mesh.Rect.Max[0] / cellSize)),
		maxY:          int(math.Ceil(navMesh.Mesh.Rect.Max[1] / cellSize)),
	}
	h.width = h.maxX - h.minX + 1
	h.height = h.maxY - h.minY + 1
	return h
}

func (h *helper) cellPointByNavMeshPoint(point *vec2.T) cellPoint {
	return cellPoint{
		int(math.Floor(point[0]/h.cellSize-float64(h.minX)) + 0.5),
		int(math.Floor(point[1]/h.cellSize-float64(h.minY)) + 0.5),
	}
}

func (h *helper) navMeshPointByCellPoint(cellPoint *cellPoint) *vec2.T {
	return &vec2.T{
		float64(cellPoint[0]+h.minX) * h.cellSize,
		float64(cellPoint[1]+h.minY) * h.cellSize,
	}
}
