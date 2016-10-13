package main

import (
	"github.com/shiwano/submarine/server/battle/server"
)

func main() {
	s := server.NewServer()
	s.Run(":5000")
}
