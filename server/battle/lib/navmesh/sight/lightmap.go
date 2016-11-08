package sight

import (
	"bytes"
	"os"

	"github.com/ugorji/go/codec"
	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// LightMap represents a map that includes lights by a point.
type LightMap struct {
	MeshVersion string     `codec:"mesh_version"`
	Helper      *helper    `codec:"helper"`
	Lights      [][]*light `codec:"lights"`
}

// GenerateLightMap creates a LightMap that has the pre-calculated lights.
func GenerateLightMap(navMesh *navmesh.NavMesh, cellSize, lightRange float64) *LightMap {
	helper := newHelper(navMesh, cellSize, lightRange)
	lm := &LightMap{
		MeshVersion: navMesh.Mesh.Version,
		Helper:      helper,
		Lights:      make([][]*light, helper.Height),
	}

	for cellY := 0; cellY < lm.Helper.Height; cellY++ {
		lm.Lights[cellY] = make([]*light, lm.Helper.Width)

		for cellX := 0; cellX < lm.Helper.Width; cellX++ {
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
func LoadLightMapFromJSONFile(fileName string) (*LightMap, error) {
	handle := new(codec.JsonHandle)
	return deserializeLightMapFromFile(fileName, handle)
}

// LoadLightMapFromMessagePackFile loads a light map from a MessagePack file.
func LoadLightMapFromMessagePackFile(fileName string) (*LightMap, error) {
	handle := new(codec.MsgpackHandle)
	return deserializeLightMapFromFile(fileName, handle)
}

func deserializeLightMapFromFile(fileName string, handle codec.Handle) (*LightMap, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	lm := new(LightMap)
	if err := codec.NewDecoder(f, handle).Decode(lm); err != nil {
		return nil, err
	}
	return lm, nil
}

// ToJSON returns JSON encoding of the light map.
func (lm *LightMap) ToJSON() ([]byte, error) {
	handle := new(codec.JsonHandle)
	return lm.serialize(handle)
}

// ToMessagePack returns MessagePack encoding of the light map.
func (lm *LightMap) ToMessagePack() ([]byte, error) {
	handle := new(codec.MsgpackHandle)
	return lm.serialize(handle)
}

func (lm *LightMap) serialize(handle codec.Handle) ([]byte, error) {
	b := new(bytes.Buffer)
	if err := codec.NewEncoder(b, handle).Encode(lm); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (lm *LightMap) lightByNavMeshPoint(point *vec2.T) *light {
	cellPoint := lm.Helper.cellPointByNavMeshPoint(point)
	if cellPoint[1] >= lm.Helper.Height ||
		cellPoint[0] >= lm.Helper.Width {
		return nil
	}
	return lm.Lights[cellPoint[1]][cellPoint[0]]
}
