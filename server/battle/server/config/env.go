package config

import (
	"flag"
	"os"
	"strings"
)

// Env is the server env.
var Env = getenv()

func getenv() string {
	submarineEnv := os.Getenv("SUBMARINE_ENV")
	if submarineEnv != "" {
		return strings.ToLower(submarineEnv)
	}

	if flag.Lookup("test.run") != nil {
		return "test"
	}
	return "development"
}
