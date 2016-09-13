package resource

import (
	"fmt"
	"lib/navmesh"
	"lib/navmesh/sight"
	"path"
	"sync"
)

// Loader loads a game resource.
var Loader = newLoader()

type loader struct {
	stageMeshes       map[int64]*navmesh.Mesh
	stagesMeshesMutex *sync.Mutex

	lightMaps      map[string]*sight.LightMap
	lightMapsMutex *sync.Mutex
}

func newLoader() *loader {
	return &loader{
		stageMeshes:       make(map[int64]*navmesh.Mesh),
		stagesMeshesMutex: new(sync.Mutex),

		lightMaps:      make(map[string]*sight.LightMap),
		lightMapsMutex: new(sync.Mutex),
	}
}

// LoadStageMesh loads the specified stage mesh.
func (l *loader) LoadStageMesh(code int64) (*navmesh.Mesh, error) {
	l.stagesMeshesMutex.Lock()
	defer l.stagesMeshesMutex.Unlock()

	if mesh, ok := l.stageMeshes[code]; ok {
		return mesh, nil
	}

	assetPath := fmt.Sprintf("Art/Stages/%03d/NavMesh.json", code)
	mesh, err := navmesh.LoadMeshFromJSONFile(path.Join(clientAssetDir, assetPath))
	if err != nil {
		return nil, err
	}
	l.stageMeshes[code] = mesh
	return mesh, nil
}

// LoadLightMap loads the specified light map.
func (l *loader) LoadLightMap(code int64, cellSize, lightRange float64) (*sight.LightMap, error) {
	l.lightMapsMutex.Lock()
	defer l.lightMapsMutex.Unlock()

	mesh, err := l.LoadStageMesh(code)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v-%v-%v", code, cellSize, lightRange)
	if lm, ok := l.lightMaps[key]; ok {
		return lm, nil
	}

	navMesh := navmesh.New(mesh)
	lm := sight.GenerateLightMap(navMesh, cellSize, lightRange)
	l.lightMaps[key] = lm
	return lm, nil
}
