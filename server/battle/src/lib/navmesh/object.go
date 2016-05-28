package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
	"math"
)

// Object represents an object in the navmesh.
type Object interface {
	ID() int64
	Position() *vec2.T
	SizeRadius() float64
	Destroy()
	IntersectWithLine(lineOrigin *vec2.T, lineVector vec2.T) *vec2.T
}

type object struct {
	id         int64
	navMesh    *NavMesh
	position   *vec2.T
	sizeRadius float64
}

// ID returns the object ID.
func (o *object) ID() int64 {
	return o.id
}

// Position returns the current position.
func (o *object) Position() *vec2.T {
	return o.position
}

// SizeRadius returns the agent size radius.
func (o *object) SizeRadius() float64 {
	return o.sizeRadius
}

// Destroy destroys self.
func (o *object) Destroy() {
	o.navMesh.DestroyObject(o.id)
}

// IntersectWithLine returns the intersection point with the given line.
func (o *object) IntersectWithLine(lineOrigin *vec2.T, lineVector vec2.T) *vec2.T {
	lineOriginFromObject := *lineOrigin
	lineOriginFromObject[0] -= o.position[0]
	lineOriginFromObject[1] -= o.position[1]
	normalizedLineVector := lineVector.Normalize()

	dotLOandLV := vec2.Dot(&lineOriginFromObject, normalizedLineVector)
	dotLOAndLO := vec2.Dot(&lineOriginFromObject, &lineOriginFromObject)

	s := dotLOandLV*dotLOandLV - dotLOAndLO + o.sizeRadius*o.sizeRadius
	if s < 0 {
		return nil
	}

	sq := math.Sqrt(s)
	t1 := -dotLOandLV - sq
	t2 := -dotLOandLV + sq
	if t1 < 0 || t2 < 0 || t1*t1 > lineVector.LengthSqr() {
		return nil
	}

	return &vec2.T{
		lineOriginFromObject[0] + normalizedLineVector[0]*t1 + o.position[0],
		lineOriginFromObject[1] + normalizedLineVector[1]*t1 + o.position[1],
	}
}
