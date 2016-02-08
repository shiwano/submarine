package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/k0kubun/pp"
)

// Log is a logrus.Logger
var Log = logrus.New()

// P is a pretty printer function.
var P = pp.Println
