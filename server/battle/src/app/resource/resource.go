package resource

import (
	"app/config"
	"path"
	"runtime"
)

var clientAssetDir string

func init() {
	_, filename, _, _ := runtime.Caller(1)
	rootDir := path.Join(path.Dir(filename), "../../../../..")

	switch config.Env {
	case "test":
		clientAssetDir = path.Join(rootDir, "server/battle/fixtures/Assets")
	case "development":
		clientAssetDir = path.Join(rootDir, "client/Assets")
	default:
		clientAssetDir = path.Join(rootDir, "client/Assets")
	}
}
