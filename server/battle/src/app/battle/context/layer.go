package context

import (
	"lib/navmesh"
)

// Navmesh layers.
const (
	LayerTeam1     = navmesh.Layer01
	LayerTeam2     = navmesh.Layer02
	LayerTeam3     = navmesh.Layer03
	LayerTeam4     = navmesh.Layer04
	LayerSubmarine = navmesh.Layer05
	LayerTorpedo   = navmesh.Layer06
)

// GetTeamLayer returns a layer mask that has the specified number.
func GetTeamLayer(n int) navmesh.LayerMask {
	switch n {
	case 1:
		return LayerTeam1
	case 2:
		return LayerTeam2
	case 3:
		return LayerTeam3
	default:
		return LayerTeam4
	}
}
