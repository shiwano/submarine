package debugger

import (
	"image"
	"sync"

	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/paint"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// Debugger represents a nevmeth debugger.
type Debugger struct {
	mu          *sync.Mutex
	event       eventSender
	screen      *image.RGBA
	screenRect  image.Rectangle
	navMeshView *navMeshView
}

func newDebugger(event eventSender) *Debugger {
	return &Debugger{
		mu:    new(sync.Mutex),
		event: event,
	}
}

// UpdateNavMesh updates nav mesh view.
func (d *Debugger) UpdateNavMesh(navMesh *navmesh.NavMesh) {
	d.mu.Lock()
	if navMesh == nil {
		d.navMeshView = nil
	} else {
		if d.navMeshView == nil {
			d.navMeshView = newNavMeshView(d.screenRect)
		}
		d.navMeshView.draw(navMesh)
	}
	d.mu.Unlock()
	d.render()
}

func (d *Debugger) render() {
	d.mu.Lock()
	if d.navMeshView != nil {
		draw.Draw(d.screen, d.screenRect, d.navMeshView.scaledMeshImage, image.ZP, draw.Src)
		draw.Draw(d.screen, d.screenRect, d.navMeshView.objectsImage, image.ZP, draw.Over)
	}
	d.mu.Unlock()
	d.event.Send(paint.Event{})
}

func (d *Debugger) setScreen(screen *image.RGBA) {
	d.mu.Lock()
	d.screen = screen
	d.screenRect = d.screen.Bounds()
	if d.navMeshView != nil {
		d.navMeshView.setScreenRect(d.screenRect)
	}
	d.mu.Unlock()
}
