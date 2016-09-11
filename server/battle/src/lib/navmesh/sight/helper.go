package sight

import (
	"lib/navmesh"
	"math"

	"github.com/ungerik/go3d/float64/vec2"
)

type cellPoint [2]int

type helper struct {
	CellSize float64 `json:"cellSize"`

	LightRange    float64 `json:"lightRange"`
	LightRangeSqr float64 `json:"lightRangeSqr"`
	LightDiameter float64 `json:"lightDiameter"`

	MinX int `json:"minX"`
	MinY int `json:"minY"`
	MaxX int `json:"maxX"`
	MaxY int `json:"maxY"`

	Width  int `json:"width"`
	Height int `json:"height"`
}

func newHelper(navMesh *navmesh.NavMesh, cellSize float64, lightRange float64) *helper {
	h := &helper{
		CellSize:      cellSize,
		LightRange:    lightRange,
		LightRangeSqr: lightRange * lightRange,
		LightDiameter: lightRange*2 + 1,
		MinX:          int(math.Floor(navMesh.Mesh.Rect.Min[0] / cellSize)),
		MinY:          int(math.Floor(navMesh.Mesh.Rect.Min[1] / cellSize)),
		MaxX:          int(math.Ceil(navMesh.Mesh.Rect.Max[0] / cellSize)),
		MaxY:          int(math.Ceil(navMesh.Mesh.Rect.Max[1] / cellSize)),
	}
	h.Width = h.MaxX - h.MinX + 1
	h.Height = h.MaxY - h.MinY + 1
	return h
}

func (h *helper) cellPointByNavMeshPoint(point *vec2.T) cellPoint {
	return cellPoint{
		int(math.Floor(point[0]/h.CellSize-float64(h.MinX)) + 0.5),
		int(math.Floor(point[1]/h.CellSize-float64(h.MinY)) + 0.5),
	}
}

func (h *helper) navMeshPointByCellPoint(cellPoint *cellPoint) *vec2.T {
	return &vec2.T{
		float64(cellPoint[0]+h.MinX) * h.CellSize,
		float64(cellPoint[1]+h.MinY) * h.CellSize,
	}
}
