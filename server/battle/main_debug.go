// +build debug

package main

import (
	"runtime"

	"github.com/shiwano/submarine/server/battle/lib/navmesh/debugger"
	"github.com/shiwano/submarine/server/battle/src"
	"github.com/shiwano/submarine/server/battle/src/debug"
)

func main() {
	debugger.Main(func(d *debugger.Debugger) {
		debug.Debugger = d
		s := server.New()
		s.Run(":5000")
	})
}

func init() {
	runtime.LockOSThread()
}
