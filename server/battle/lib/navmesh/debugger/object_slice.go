package debugger

import (
	"sort"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

type objectSlice []navmesh.Object

func newObjectSlice(objects map[int64]navmesh.Object) objectSlice {
	os := make(objectSlice, len(objects))
	i := 0
	for _, o := range objects {
		os[i] = o
		i++
	}
	return os
}

func (os objectSlice) Len() int           { return len(os) }
func (os objectSlice) Swap(i, j int)      { os[i], os[j] = os[j], os[i] }
func (os objectSlice) Less(i, j int) bool { return os[i].ID() < os[j].ID() }

func (os objectSlice) sort() {
	sort.Sort(os)
}
