package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/shiwano/submarine/server/battle/lib/typhenapi"
	configAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/configuration"
)

// Config is the loaded server config.
var Config *configAPI.Server

func newServerConfig() (*configAPI.Server, error) {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Join(filepath.Dir(filename), "../../")
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	var path string
	if Env == "test" {
		path = filepath.Join(dir, "config.test.json")
	} else {
		path = filepath.Join(dir, "config.json")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	serializer := new(typhenapi.JSONSerializer)
	c := new(configAPI.Server)
	err = serializer.Deserialize(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	serverConfig, err := newServerConfig()
	if err != nil {
		panic(err)
	}
	Config = serverConfig
}
