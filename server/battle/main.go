// +build !debug

package main

import (
	"github.com/shiwano/submarine/server/battle/server"
)

func main() {
	s := server.New()
	s.Run(":5000")
}
