package resource

import (
	"app/config"
	"lib/navmesh"
	"path"
	"runtime"
)

var clientAssetDir string

// BattleMap represents a battle map resource.
type BattleMap struct {
	Code    int64
	NavMesh *navmesh.NavMesh
}

func init() {
	_, filename, _, _ := runtime.Caller(1)
	rootDir := path.Join(path.Dir(filename), "../../../../..")

	switch config.Env {
	case "test":
		clientAssetDir = path.Join(rootDir, "server/battle/fixtures/Assets")
	case "development":
		clientAssetDir = path.Join(rootDir, "client/Assets")
	}
}
