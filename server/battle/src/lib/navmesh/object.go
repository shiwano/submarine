package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Object represents an object in the navmesh.
type Object interface {
	ID() int64
	Position() *vec2.T
	Size() float64
}
