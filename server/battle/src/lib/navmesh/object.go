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

	callCollideHandler(obj Object, point vec2.T)
	SetCollideHandler(handler func(obj Object, point vec2.T))

	IntersectWithLine(lineOrigin, lineDir, lineVector *vec2.T) (vec2.T, bool)
}

type object struct {
	id             int64
	navMesh        *NavMesh
	position       *vec2.T
	sizeRadius     float64
	collideHandler func(obj Object, point vec2.T)
}

func (o *object) ID() int64 {
	return o.id
}

func (o *object) Position() *vec2.T {
	return o.position
}

func (o *object) SizeRadius() float64 {
	return o.sizeRadius
}

func (o *object) Destroy() {
	o.navMesh.DestroyObject(o.id)
}

func (o *object) callCollideHandler(obj Object, point vec2.T) {
	if o.collideHandler != nil {
		o.collideHandler(obj, point)
	}
}

func (o *object) SetCollideHandler(handler func(obj Object, point vec2.T)) {
	o.collideHandler = handler
}

func (o *object) IntersectWithLine(lineOrigin, lineDir, lineVector *vec2.T) (vec2.T, bool) {
	lineOriginFromObject := vec2.Sub(lineOrigin, o.position)
	normalizedLineVector := lineDir.Normalized()

	dotLOandLV := vec2.Dot(&lineOriginFromObject, &normalizedLineVector)
	dotLOAndLO := vec2.Dot(&lineOriginFromObject, &lineOriginFromObject)

	s := dotLOandLV*dotLOandLV - dotLOAndLO + o.sizeRadius*o.sizeRadius
	if s < 0 {
		return vec2.Zero, false
	}

	sq := math.Sqrt(s)
	t1 := -dotLOandLV - sq
	t2 := -dotLOandLV + sq
	if t1 < 0 || t2 < 0 || t1*t1 > lineVector.LengthSqr() {
		return vec2.Zero, false
	}

	resultPoint := normalizedLineVector.Scaled(t1)
	resultPoint.Add(lineOrigin)
	return resultPoint, true
}
