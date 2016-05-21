package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Edge repsents a edge on the mesh.
type Edge [2]*vec2.T

func (e Edge) intersect(p1, p2 *vec2.T) bool {
	v1 := e[0]
	v2 := e[1]
	return ((p1[0]-p2[0])*(v1[1]-p1[1])+(p1[1]-p2[1])*(p1[0]-v1[0]))*
		((p1[0]-p2[0])*(v2[1]-p1[1])+(p1[1]-p2[1])*(p1[0]-v2[0])) < 0 &&
		((v1[0]-v2[0])*(p1[1]-v1[1])+(v1[1]-v2[1])*(v1[0]-p1[0]))*
			((v1[0]-v2[0])*(p2[1]-v1[1])+(v1[1]-v2[1])*(v1[0]-p2[0])) < 0
}
