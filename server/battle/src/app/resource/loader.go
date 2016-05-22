package resource

import (
	"fmt"
	"lib/navmesh"
	"path"
	"sync"
)

// Loader loads a game resource.
var Loader = newLoader()

type loader struct {
	stageMeshes       map[int64]*navmesh.Mesh
	stagesMeshesMutex *sync.Mutex
}

func newLoader() *loader {
	return &loader{
		stageMeshes:       make(map[int64]*navmesh.Mesh),
		stagesMeshesMutex: new(sync.Mutex),
	}
}

// LoadStageMesh loads the specified stage mesh.
func (l *loader) LoadStageMesh(code int64) (*navmesh.Mesh, error) {
	l.stagesMeshesMutex.Lock()
	defer l.stagesMeshesMutex.Unlock()

	if mesh, ok := l.stageMeshes[code]; ok {
		return mesh, nil
	}

	assetPath := fmt.Sprintf("Art/Maps/%03d/NavMesh.json", code)
	mesh, err := navmesh.LoadMeshFromJSONFile(path.Join(clientAssetDir, assetPath))
	if err != nil {
		return nil, err
	}
	l.stageMeshes[code] = mesh
	return mesh, nil
}
