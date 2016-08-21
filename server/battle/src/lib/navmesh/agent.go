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
func (a *Agent) Move(position *vec2.T, ignoredLayer LayerMask) {
	vec := vec2.Sub(position, a.position)

	if hitInfo, ok := a.navMesh.Raycast(a.position, &vec, ignoredLayer); ok {
		a.position = &hitInfo.point

		a.callCollideHandler(hitInfo.obj, hitInfo.point)
		if hitInfo.obj != nil {
			hitInfo.obj.callCollideHandler(a, hitInfo.point)
		}
		return
	}
	a.position = position
}
