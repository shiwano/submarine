package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
	"math"
)

func equalVectors(v1, v2 *vec2.T) bool {
	return v1[0] == v2[0] && v1[1] == v2[1]
}

func calculateOctileDistance(from, to *vec2.T) float64 {
	dx := math.Abs(from[0] - to[0])
	dy := math.Abs(from[1] - to[1])
	return (dx + dy) + (math.Sqrt2-2)*math.Min(dx, dy)
}

func calculateVectorLengthSqr(src, dest *vec2.T) float64 {
	vector := vec2.Sub(dest, src)
	return vector.LengthSqr()
}
