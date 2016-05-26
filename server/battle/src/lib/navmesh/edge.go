package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Edge represents a edge on the mesh.
type Edge [2]*vec2.T

func (e Edge) vector() vec2.T {
	return vec2.Sub(e[1], e[0])
}

func (e Edge) cross(a, b *vec2.T) float64 {
	return a[1]*b[0] - a[0]*b[1]
}

func (e Edge) intersect(p1, p2 *vec2.T) *vec2.T {
	v1 := e.vector()
	v2 := vec2.Sub(p2, p1)

	crossV1andV2 := e.cross(&v1, &v2)
	if crossV1andV2 == 0 {
		return nil
	}

	v := vec2.Sub(p1, e[0])
	crossVandV1 := e.cross(&v, &v1)
	t2 := crossVandV1 / crossV1andV2
	if t2 < 0 || t2 > 1 {
		return nil
	}

	crossVandV2 := e.cross(&v, &v2)
	t1 := crossVandV2 / crossV1andV2
	if t1 < 0 || t1 > 1 {
		return nil
	}

	intersectionPoint := v1.Scaled(t1)
	intersectionPoint.Add(e[0])
	return &intersectionPoint
}
