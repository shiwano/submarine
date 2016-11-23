package debugger

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/paint"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// Debugger represents a nevmeth debugger.
type Debugger struct {
	mu          *sync.Mutex
	event       screen.EventDeque
	screen      screen.Buffer
	navMeshView *navMeshView
}

func newDebugger(event screen.EventDeque) *Debugger {
	return &Debugger{
		mu:    new(sync.Mutex),
		event: event,
	}
}

// UpdateNavMesh updates nav mesh view.
func (d *Debugger) UpdateNavMesh(navMesh *navmesh.NavMesh) {
	d.mu.Lock()
	d.drawNavMesh(navMesh)
	d.drawScreen()
	d.mu.Unlock()
	d.event.Send(paint.Event{})
}

func (d *Debugger) close() {
	d.mu.Lock()
	d.screen.Release()
	d.screen = nil
	d.mu.Unlock()
}

func (d *Debugger) setScreen(s screen.Buffer) {
	d.mu.Lock()
	if d.screen != nil {
		d.screen.Release()
	}
	d.screen = s
	if d.navMeshView != nil {
		d.navMeshView.setScreenRect(d.screen.Bounds())
	}
	d.drawScreen()
	d.mu.Unlock()
	d.event.Send(paint.Event{})
}

func (d *Debugger) uploadBuffer(uploader screen.Uploader) {
	d.mu.Lock()
	uploader.Upload(image.ZP, d.screen, d.screen.Bounds())
	d.mu.Unlock()
}

func (d *Debugger) drawNavMesh(navMesh *navmesh.NavMesh) {
	if navMesh == nil {
		d.navMeshView = nil
		return
	}
	if d.navMeshView == nil {
		d.navMeshView = newNavMeshView(d.screen.Bounds())
	}
	d.navMeshView.draw(navMesh)
}

func (d *Debugger) drawScreen() {
	if d.navMeshView == nil {
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), image.Transparent, image.ZP, draw.Src)
	} else {
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), d.navMeshView.meshImage, image.ZP, draw.Src)
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), d.navMeshView.objectsImage, image.ZP, draw.Over)
	}
}
