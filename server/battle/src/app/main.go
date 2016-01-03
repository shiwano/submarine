package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

// Log is a logrus.Logger
var Log = logrus.New()

var p = fmt.Println

func main() {
	server := NewServer()
	defer server.Close()
	server.Run(":5000")
}
