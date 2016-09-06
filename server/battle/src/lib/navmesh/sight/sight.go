package sight

import (
	"github.com/ungerik/go3d/float64/vec2"
	"strings"
)

// Sight represents the sight of the battle.
type Sight struct {
	cells     [][]bool
	lightMap  *LightMap
	putLights map[*light]struct{}
}

// New creates a sight.
func New(lightMap *LightMap) *Sight {
	s := &Sight{
		lightMap: lightMap,
	}
	s.Clear()
	return s
}

// Clear clears all lit point informations and the lights.
func (s *Sight) Clear() {
	s.putLights = make(map[*light]struct{})

	s.cells = make([][]bool, s.lightMap.Helper.height)
	for cellY := 0; cellY < s.lightMap.Helper.height; cellY++ {
		s.cells[cellY] = make([]bool, s.lightMap.Helper.width)
	}
}

// PutLight puts a light to the specified point.
func (s *Sight) PutLight(point *vec2.T) {
	light := s.lightMap.lightByNavMeshPoint(point)
	if light == nil {
		return
	}
	if _, ok := s.putLights[light]; !ok {
		s.putLights[light] = struct{}{}
		for _, cellPoint := range light.LitPoints {
			s.cells[cellPoint[1]][cellPoint[0]] = true
		}
	}
}

// IsLitPoint determines whether the specified point is lit.
func (s *Sight) IsLitPoint(point *vec2.T) bool {
	cellPoint := s.lightMap.Helper.cellPointByNavMeshPoint(point)
	if cellPoint[1] >= s.lightMap.Helper.height ||
		cellPoint[0] >= s.lightMap.Helper.width {
		return false
	}
	return s.cells[cellPoint[1]][cellPoint[0]]
}

// DebugString returns a string that represents the lit points for debug.
func (s *Sight) DebugString() string {
	strs := make([]string, len(s.cells))
	for y, cellsX := range s.cells {
		runes := make([]rune, len(s.cells[y]))
		for x, cell := range cellsX {
			if cell {
				runes[x] = 'X'
			} else {
				runes[x] = '_'
			}
		}
		strs[len(s.cells)-y-1] = string(runes)
	}
	return strings.Join(strs, "\n")
}
