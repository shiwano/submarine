package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Agent represents a movable object in the navmesh.
type Agent struct {
	*object
	currentTriangle *Triangle
}

// Move sets the specified position to the agent if the position is out of the mesh.
func (a *Agent) Move(position *vec2.T) bool {
	diff := vec2.Sub(position, a.position)
	positionWithSizeRadius := diff.Normalize().Scale(a.sizeRadius).Add(position)

	if a.currentTriangle != nil && a.currentTriangle.containsPoint(positionWithSizeRadius) {
		a.position = position
		return true
	} else if a.currentTriangle = a.navMesh.Mesh.findTriangleByPoint(positionWithSizeRadius); a.currentTriangle != nil {
		a.position = position
		return true
	}
	return false
}
