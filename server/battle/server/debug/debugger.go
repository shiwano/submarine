package debug

import (
	"sort"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/debugger/interface"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Debugger is the static variable that will be set by main function for debug.
var Debugger debugger.Debugger

// SortedSights returns sorted sight slice by layer.
func SortedSights(sightsByLayer map[navmesh.LayerMask]*sight.Sight) []*sight.Sight {
	sightsKeys := make([]int, len(sightsByLayer))
	i := 0
	for k := range sightsByLayer {
		sightsKeys[i] = int(k)
		i++
	}
	sort.Ints(sightsKeys)

	sights := make([]*sight.Sight, len(sightsByLayer))
	for i, k := range sightsKeys {
		sights[i] = sightsByLayer[navmesh.LayerMask(k)]
	}
	return sights
}
