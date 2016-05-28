package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Object represents an object in the navmesh.
type Object interface {
	ID() int64
	Position() *vec2.T
	SizeRadius() float64
	Destroy()
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
	o.navMesh.DestroyObject(o.ID())
}
