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

func (e *edge) containsPoint(point *vec2.T) bool {
	pointVec := vec2.Sub(point, e.points[0])
	vectorLengthSqr := e.vector.LengthSqr()
	pointVecLengthSqr := pointVec.LengthSqr()

	if vectorLengthSqr < pointVecLengthSqr {
		return false
	}
	dot := vec2.Dot(e.vector, &pointVec)
	if !equalFloats(dot*dot, vectorLengthSqr*pointVecLengthSqr) {
		return false
	}
	return true
}

func (e *edge) intersectWithLineSeg(lineOrigin, lineVec *vec2.T) (vec2.T, bool) {
	crossEVandLV := cross(e.vector, lineVec)
	if equalFloats(crossEVandLV, 0) {
		return vec2.Zero, false
	}

	v := vec2.Sub(lineOrigin, e.points[0])
	crossVandEV := cross(&v, e.vector)
	t2 := crossVandEV / crossEVandLV
	if t2 < 0 || t2 > 1 {
		return vec2.Zero, false
	}

	crossVandLV := cross(&v, lineVec)
	t1 := crossVandLV / crossEVandLV
	if t1 < 0 || t1 > 1 {
		return vec2.Zero, false
	}

	resultPoint := e.vector.Scaled(t1)
	resultPoint.Add(e.points[0])
	return resultPoint, true
}
