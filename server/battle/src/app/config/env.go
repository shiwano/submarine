package config

import (
	"os"
	"strings"
)

// Env is the server env.
var Env = getEnv()

func getEnv() string {
	env := strings.ToLower(os.Getenv("SUBMARINE_ENV"))
	if env == "" {
		env = "development"
	}
	return env
}
