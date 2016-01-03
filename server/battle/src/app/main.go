package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/k0kubun/pp"
)

// Log is a logrus.Logger
var Log = logrus.New()

var p = pp.Println

func main() {
	server := NewServer()
	server.Run(":5000")
}
