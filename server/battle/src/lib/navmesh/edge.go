package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

type edge struct {
	points [2]*vec2.T
	vector *vec2.T
}

func newEdge(a, b *vec2.T) *edge {
	vector := vec2.Sub(b, a)
	return &edge{
		points: [2]*vec2.T{a, b},
		vector: &vector,
	}
}

func (e *edge) cross(a, b *vec2.T) float64 {
	return a[1]*b[0] - a[0]*b[1]
}

func (e *edge) intersectWithLine(lineOrigin, lineVector *vec2.T) (vec2.T, bool) {
	crossEVandLV := e.cross(e.vector, lineVector)
	if crossEVandLV == 0 {
		return vec2.Zero, false
	}

	v := vec2.Sub(lineOrigin, e.points[0])
	crossVandEV := e.cross(&v, e.vector)
	t2 := crossVandEV / crossEVandLV
	if t2 < 0 || t2 > 1 {
		return vec2.Zero, false
	}

	crossVandLV := e.cross(&v, lineVector)
	t1 := crossVandLV / crossEVandLV
	if t1 < 0 || t1 > 1 {
		return vec2.Zero, false
	}

	resultPoint := e.vector.Scaled(t1)
	resultPoint.Add(e.points[0])
	return resultPoint, true
}
