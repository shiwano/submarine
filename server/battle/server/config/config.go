package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Config is the loaded server config.
var Config *ServerConfig

// ServerConfig represents the game server config.
type ServerConfig struct {
	APIServerBaseURI    string `yaml:"api_server_base_uri"`
	BattleServerBaseURI string `yaml:"battle_server_base_uri"`
}

func newServerConfig() (*ServerConfig, error) {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Join(filepath.Dir(filename), "../../")
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	var path string
	if Env == "test" {
		path = filepath.Join(dir, "config.example.yml")
	} else {
		path = filepath.Join(dir, "config.yml")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := new(ServerConfig)
	err = yaml.Unmarshal(data, c)
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
