package sight

import (
	"encoding/json"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"os"
)

// LightMap represents a map that includes lights by a point.
type LightMap struct {
	Helper *helper    `json:"helper"`
	Lights [][]*light `json:"lights"`
}

// GenerateLightMap creates a LightMap that has the pre-calculated lights.
func GenerateLightMap(navMesh *navmesh.NavMesh, cellSize float64, lightRange float64) *LightMap {
	helper := newHelper(navMesh, cellSize, lightRange)
	lm := &LightMap{
		Helper: helper,
		Lights: make([][]*light, helper.height),
	}

	for cellY := 0; cellY < lm.Helper.height; cellY++ {
		lm.Lights[cellY] = make([]*light, lm.Helper.width)

		for cellX := 0; cellX < lm.Helper.width; cellX++ {
			cellPoint := &cellPoint{cellX, cellY}
			point := lm.Helper.navMeshPointByCellPoint(cellPoint)

			if light := newLight(navMesh, lm.Helper, point); light.isLighting() {
				lm.Lights[cellY][cellX] = light
			}
		}
	}
	return lm
}

// LoadLightMapFromJSONFile loads a light map from a JSON file.
func LoadLightMapFromJSONFile(jsonPath string) (*LightMap, error) {
	f, err := os.Open(jsonPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	lm := new(LightMap)
	if err := json.NewDecoder(f).Decode(lm); err != nil {
		return nil, err
	}
	return lm, nil
}

// ToJSON returns JSON encoding of the light map.
func (lm *LightMap) ToJSON() ([]byte, error) {
	return json.Marshal(lm)
}

func (lm *LightMap) lightByNavMeshPoint(point *vec2.T) *light {
	cellPoint := lm.Helper.cellPointByNavMeshPoint(point)
	if cellPoint[1] >= lm.Helper.height ||
		cellPoint[0] >= lm.Helper.width {
		return nil
	}
	return lm.Lights[cellPoint[1]][cellPoint[0]]
}
