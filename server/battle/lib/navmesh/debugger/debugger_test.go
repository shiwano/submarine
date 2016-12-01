package debugger

import (
	"testing"

	iface "github.com/shiwano/submarine/server/battle/lib/navmesh/debugger/interface"
)

func TestDebugger(t *testing.T) {
	if _, ok := interface{}((*Debugger)(nil)).(iface.Debugger); !ok {
		t.Error("should implements Debugger interface")
	}
}
