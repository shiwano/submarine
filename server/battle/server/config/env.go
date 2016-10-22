package config

import (
	"flag"
	"os"
	"strings"
)

// server env.
var (
	Env            = getenv()
	EnvDevelopment = "development"
	EnvTest        = "test"
)

func getenv() string {
	submarineEnv := os.Getenv("SUBMARINE_ENV")
	if submarineEnv != "" {
		return strings.ToLower(submarineEnv)
	}

	if flag.Lookup("test.run") != nil {
		return EnvTest
	}
	return EnvDevelopment
}
