package component

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// Visibility represents visibility of an actor, it manages visibility in duplicate.
type Visibility struct {
	countsByTeam  map[navmesh.LayerMask]int
	ChangeHandler func(navmesh.LayerMask)
}

// NewVisibility creates a visibility.
func NewVisibility() *Visibility {
	return &Visibility{
		countsByTeam: make(map[navmesh.LayerMask]int),
	}
}

// IsVisibleFrom determines whether the actor is visible from the specified layer.
func (v *Visibility) IsVisibleFrom(layer navmesh.LayerMask) bool {
	return v.countsByTeam[layer] > 0
}

// Set the actor visibility.
func (v *Visibility) Set(layer navmesh.LayerMask, isVisible bool) {
	previousVisibility := v.IsVisibleFrom(layer)

	if isVisible {
		v.countsByTeam[layer]++
	} else if v.countsByTeam[layer] > 0 {
		v.countsByTeam[layer]--
	}

	if v.ChangeHandler != nil && previousVisibility != v.IsVisibleFrom(layer) {
		v.ChangeHandler(layer)
	}
}
