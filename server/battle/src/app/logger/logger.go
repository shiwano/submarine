package logger

import (
	"app/config"
	"github.com/Sirupsen/logrus"
	"github.com/k0kubun/pp"
)

// Log is a logrus.Logger
var Log = logrus.New()

// P is a pretty printer function.
var P = pp.Println

func init() {
	switch config.Env {
	case "test":
		Log.Level = logrus.PanicLevel
	case "development":
		Log.Level = logrus.DebugLevel
	}
}
