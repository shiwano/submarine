package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
	"math"
)

// FloatEpsilon is a tolerance for comparing floating point values.
var FloatEpsilon = 0.000001

func equalVectors(v1, v2 *vec2.T) bool {
	return equalFloats(v1[0], v2[0]) && equalFloats(v1[1], v2[1])
}

func equalFloats(a, b float64) bool {
	return math.Abs(a-b) < FloatEpsilon
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

func cross(a, b *vec2.T) float64 {
	return a[1]*b[0] - a[0]*b[1]
}
