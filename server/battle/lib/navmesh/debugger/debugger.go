package debugger

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/paint"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/lib/navmesh/sight"
)

// Debugger represents a nevmeth debugger.
type Debugger struct {
	mu          *sync.Mutex
	event       screen.EventDeque
	screen      screen.Buffer
	navMeshView *navMeshView
	sightViews  []*sightView
}

func newDebugger(event screen.EventDeque) *Debugger {
	return &Debugger{
		mu:    new(sync.Mutex),
		event: event,
	}
}

// Update updates the debugger window.
func (d *Debugger) Update(navMesh *navmesh.NavMesh, sights []*sight.Sight) {
	d.mu.Lock()
	d.drawNavMesh(navMesh)
	if len(sights) > 0 && navMesh != nil {
		d.drawSights(sights, navMesh.Mesh)
	}
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
	for _, sightView := range d.sightViews {
		sightView.setScreenRect(d.screen.Bounds())
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

func (d *Debugger) drawSights(sights []*sight.Sight, mesh *navmesh.Mesh) {
	for i, s := range sights {
		if len(d.sightViews) <= i {
			sv := newSightView(d.screen.Bounds())
			d.sightViews = append(d.sightViews, sv)
		}
		sv := d.sightViews[i]
		sv.draw(s, mesh)
	}
}

func (d *Debugger) drawScreen() {
	if d.navMeshView == nil {
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), image.Transparent, image.ZP, draw.Src)
	} else {
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), d.navMeshView.meshImage, image.ZP, draw.Src)
		draw.Draw(d.screen.RGBA(), d.screen.Bounds(), d.navMeshView.objectsImage, image.ZP, draw.Over)

		if len(d.sightViews) > 0 {
			draw.Draw(d.screen.RGBA(), d.screen.Bounds(), d.sightViews[0].sightImage, image.ZP, draw.Over)
		}
	}
}
