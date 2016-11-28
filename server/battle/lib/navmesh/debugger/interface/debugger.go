package debugger

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Debugger is the Debugger interface. See also implemented debugger package.
type Debugger interface {
	Update(navMesh *navmesh.NavMesh, sights []*sight.Sight)
}
