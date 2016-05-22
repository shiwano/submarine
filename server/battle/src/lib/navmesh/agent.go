package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Agent represents a movable object in the navmesh.
type Agent struct {
	id              int64
	navMesh         *NavMesh
	currentTriangle *Triangle
	position        *vec2.T
	size            float64
}

// ID returns the object ID.
func (a *Agent) ID() int64 {
	return a.id
}

// Position returns the current position.
func (a *Agent) Position() *vec2.T {
	return a.position
}

// Size returns the agent size.
func (a *Agent) Size() float64 {
	return a.size
}

// Destroy destroys self.
func (a *Agent) Destroy() {
	a.navMesh.DestroyObject(a.id)
}

// Move sets the specified position to the agent if the position is out of the mesh.
func (a *Agent) Move(position *vec2.T) bool {
	if a.currentTriangle != nil && a.currentTriangle.containsPoint(position) {
		a.position = position
		return true
	} else if a.currentTriangle = a.navMesh.Mesh.findTriangleByPoint(position); a.currentTriangle != nil {
		a.position = position
		return true
	}
	return false
}
