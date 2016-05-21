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
	battleMapMeshes      map[int64]*navmesh.Mesh
	battleMapMeshesMutex *sync.Mutex
}

func newLoader() *loader {
	return &loader{
		battleMapMeshes:      make(map[int64]*navmesh.Mesh),
		battleMapMeshesMutex: new(sync.Mutex),
	}
}

// LoadBattleMap loads the specified battle map.
func (l *loader) LoadBattleMap(code int64) (*BattleMap, error) {
	mesh, err := l.getOrLoadBattleMapMesh(code)
	if err != nil {
		return nil, err
	}
	battleMap := &BattleMap{
		Code:    code,
		NavMesh: navmesh.New(mesh),
	}
	return battleMap, nil
}

func (l *loader) getOrLoadBattleMapMesh(code int64) (*navmesh.Mesh, error) {
	l.battleMapMeshesMutex.Lock()
	defer l.battleMapMeshesMutex.Unlock()

	if mesh, ok := l.battleMapMeshes[code]; ok {
		return mesh, nil
	}

	assetPath := fmt.Sprintf("Art/Maps/%03d/NavMesh.json", code)
	mesh, err := navmesh.LoadMeshFromJSONFile(path.Join(clientAssetDir, assetPath))
	if err != nil {
		return nil, err
	}
	l.battleMapMeshes[code] = mesh
	return mesh, nil
}
