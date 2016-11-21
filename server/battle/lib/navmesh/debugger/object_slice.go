package debugger

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

type objectSlice []navmesh.Object

func (o objectSlice) Len() int {
	return len(o)
}

func (o objectSlice) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o objectSlice) Less(i, j int) bool {
	return o[i].ID() < o[j].ID()
}
