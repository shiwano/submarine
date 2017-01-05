package resource

import (
	"fmt"
	"path"
	"sync"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Loader loads a game resource.
var Loader = newLoader()

type loader struct {
	stageMeshes       map[int64]*navmesh.Mesh
	stagesMeshesMutex *sync.Mutex

	lightMaps      map[int64]*sight.LightMap
	lightMapsMutex *sync.Mutex
}

func newLoader() *loader {
	return &loader{
		stageMeshes:       make(map[int64]*navmesh.Mesh),
		stagesMeshesMutex: new(sync.Mutex),

		lightMaps:      make(map[int64]*sight.LightMap),
		lightMapsMutex: new(sync.Mutex),
	}
}

// LoadMesh loads the specified stage mesh.
func (l *loader) LoadMesh(code int64) (*navmesh.Mesh, error) {
	l.stagesMeshesMutex.Lock()
	defer l.stagesMeshesMutex.Unlock()

	if mesh, ok := l.stageMeshes[code]; ok {
		return mesh, nil
	}

	assetPath := fmt.Sprintf("stages/%03d/mesh.json", code)
	mesh, err := navmesh.LoadMeshFromJSONFile(path.Join(clientAssetDir, assetPath))
	if err != nil {
		return nil, err
	}
	l.stageMeshes[code] = mesh
	return mesh, nil
}

// LoadLightMap loads the specified light map.
func (l *loader) LoadLightMap(code int64) (*sight.LightMap, error) {
	l.lightMapsMutex.Lock()
	defer l.lightMapsMutex.Unlock()

	if lm, ok := l.lightMaps[code]; ok {
		return lm, nil
	}

	assetPath := fmt.Sprintf("stages/%03d/lightmap.mpac", code)
	lm, err := sight.LoadLightMapFromMessagePackFile(path.Join(clientAssetDir, assetPath))
	if err != nil {
		return nil, err
	}
	l.lightMaps[code] = lm
	return lm, nil
}
