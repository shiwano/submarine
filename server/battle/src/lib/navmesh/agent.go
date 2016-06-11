package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Agent represents a movable object in the navmesh.
type Agent struct {
	*object
}

// Move sets the specified position to the agent if the position is out of the mesh.
func (a *Agent) Move(position *vec2.T) bool {
	vector := vec2.Sub(position, a.position)
	sizeRadiusVector := vector.Normalized()
	sizeRadiusVector.Scale(a.sizeRadius)
	vector.Add(&sizeRadiusVector)

	if intersectionObj, intersectionPoint, ok := a.navMesh.Raycast(a.position, &vector, filterTrigger); ok {
		a.position = &intersectionPoint
		a.position.Sub(&sizeRadiusVector)

		a.callCollideHandler(intersectionObj, intersectionPoint)
		if intersectionObj != nil {
			intersectionObj.callCollideHandler(a, intersectionPoint)
		}
		return false
	}
	a.position = position
	return true
}
