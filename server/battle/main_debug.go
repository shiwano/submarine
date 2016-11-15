// +build debug

package main

import (
	"runtime"

	"github.com/shiwano/submarine/server/battle/lib/navmesh/debugger"
	"github.com/shiwano/submarine/server/battle/server"
	"github.com/shiwano/submarine/server/battle/server/debug"
)

func main() {
	debugger.Main(func(debugger *debugger.Debugger) {
		debug.Debugger = debugger
		s := server.New()
		s.Run(":5000")
	})
}

func init() {
	runtime.LockOSThread()
}
