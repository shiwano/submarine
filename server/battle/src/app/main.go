package main

import (
	"github.com/Sirupsen/logrus"
)

// Log is a logrus.Logger
var Log = logrus.New()

func main() {
	server := NewServer()
	defer server.Close()
	server.Run(":5000")
}
