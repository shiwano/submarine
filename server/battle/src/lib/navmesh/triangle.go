package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Triangle represents a triangle on a mesh.
type Triangle struct {
	Vertices [3]*vec2.T
}

func newTriangle(vertex1, vertex2, vertex3 *vec2.T) *Triangle {
	return &Triangle{
		Vertices: [3]*vec2.T{vertex1, vertex2, vertex3},
	}
}

// via http://www.blackpawn.com/texts/pointinpoly/default.html
func (t *Triangle) containsPoint(point *vec2.T) bool {
	v0 := vec2.Sub(t.Vertices[2], t.Vertices[0])
	v1 := vec2.Sub(t.Vertices[1], t.Vertices[0])
	v2 := vec2.Sub(point, t.Vertices[0])

	dot00 := vec2.Dot(&v0, &v0)
	dot01 := vec2.Dot(&v0, &v1)
	dot02 := vec2.Dot(&v0, &v2)
	dot11 := vec2.Dot(&v1, &v1)
	dot12 := vec2.Dot(&v1, &v2)

	invDenom := 1 / (dot00*dot11 - dot01*dot01)
	u := (dot11*dot02 - dot01*dot12) * invDenom
	v := (dot00*dot12 - dot01*dot02) * invDenom

	return (u >= 0) && (v >= 0) && (u+v <= 1+FloatEpsilon)
}

func (t *Triangle) hasVertex(vertex *vec2.T) bool {
	for _, v := range t.Vertices {
		if v == vertex {
			return true
		}
	}
	return false
}

func (t *Triangle) vertexIndex(vertex *vec2.T) (int, bool) {
	for i, v := range t.Vertices {
		if v == vertex {
			return i, true
		}
	}
	return -1, false
}
