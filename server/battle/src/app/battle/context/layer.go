package context

import (
	"lib/navmesh"
)

// Navmesh layers constants.
const (
	LayerTeam1     = navmesh.Layer01
	LayerTeam2     = navmesh.Layer02
	LayerTeam3     = navmesh.Layer03
	LayerTeam4     = navmesh.Layer04
	LayerSubmarine = navmesh.Layer05
	LayerTorpedo   = navmesh.Layer06
)

// Navmesh layers variables.
var (
	TeamLayers = [4]navmesh.LayerMask{
		LayerTeam1,
		LayerTeam2,
		LayerTeam3,
		LayerTeam4,
	}
)

// GetTeamLayer returns a layer mask that has the specified number.
func GetTeamLayer(n int) navmesh.LayerMask {
	return TeamLayers[n-1]
}
