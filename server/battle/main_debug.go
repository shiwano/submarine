// +build debug

package main

import (
	"runtime"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/debugger"
	"github.com/shiwano/submarine/server/battle/server"
	"github.com/shiwano/submarine/server/battle/server/debug"
)

func main() {
	debugger.Main(func(d navmesh.Debugger) {
		debug.Debugger = d
		s := server.New()
		s.Run(":5000")
	})
}

func init() {
	runtime.LockOSThread()
}
