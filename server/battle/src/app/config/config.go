package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Config is the loaded server config.
var Config = newServerConfig()

// ServerConfig represents the game server config.
type ServerConfig struct {
	APIServerBaseURI    string `yaml:"api_server_base_uri"`
	BattleServerBaseURI string `yaml:"battle_server_base_uri"`
}

func newServerConfig() *ServerConfig {
	_, filename, _, _ := runtime.Caller(1)
	dir, err := filepath.Abs(filepath.Dir(filename) + "/../../../../../")
	if err != nil {
		panic(err)
	}

	var path string
	if Env == "test" {
		path = filepath.Join(dir, "config.example.yml")
	} else {
		path = filepath.Join(dir, "config."+Env+".yml")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	c := new(struct {
		Server ServerConfig `yaml:"server"`
	})
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return &c.Server
}
