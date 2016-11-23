package sight

import (
	"strings"

	"github.com/ungerik/go3d/float64/vec2"
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

// LitPoints returns lit points on the navmesh.
func (s *Sight) LitPoints() []vec2.T {
	var points []vec2.T
	for y, cellsY := range s.cells {
		for x, isLit := range cellsY {
			if isLit {
				cellPoint := &cellPoint{x, y}
				p := s.lightMap.Helper.navMeshPointByCellPoint(cellPoint)
				points = append(points, p)
			}
		}
	}
	return points
}

// Clear clears all lit point informations and the lights.
func (s *Sight) Clear() {
	s.putLights = make(map[*light]struct{})

	s.cells = make([][]bool, s.lightMap.Helper.Height)
	for cellY := 0; cellY < s.lightMap.Helper.Height; cellY++ {
		s.cells[cellY] = make([]bool, s.lightMap.Helper.Width)
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
	if cellPoint[1] >= s.lightMap.Helper.Height ||
		cellPoint[0] >= s.lightMap.Helper.Width {
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
