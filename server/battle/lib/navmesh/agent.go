package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// Agent represents a movable object in the navmesh.
type Agent struct {
	*object
}

// Warp sets the actor position.
func (a *Agent) Warp(position *vec2.T) {
	a.position = position
}

// Move sets the actor position. If the specified position is collided with
// the mesh or objects, this method sets the collided position.
func (a *Agent) Move(position *vec2.T, ignoredLayer LayerMask) *RaycastHitInfo {
	vec := vec2.Sub(position, a.position)

	if hitInfo := a.navMesh.Raycast(a.position, &vec, ignoredLayer); hitInfo != nil {
		a.position = &hitInfo.Point
		return hitInfo
	}
	a.position = position
	return nil
}
