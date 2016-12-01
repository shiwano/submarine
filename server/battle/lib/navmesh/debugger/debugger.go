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

// Debugger represents a nevmesh debugger.
type Debugger struct {
	mu     *sync.Mutex
	event  screen.EventDeque
	screen screen.Buffer

	navMeshView *navMeshView
	sightViews  []*sightView
}

func newDebugger(event screen.EventDeque) *Debugger {
	return &Debugger{
		mu:    new(sync.Mutex),
		event: event,
	}
}

// Update the debugger. This function will be executed goroutine-safe.
// If you want to clear the screen, give nil to all parameters.
func (d *Debugger) Update(navMesh *navmesh.NavMesh, sights []*sight.Sight) {
	d.mu.Lock()
	draw.Draw(d.screen.RGBA(), d.screen.Bounds(), image.Transparent, image.ZP, draw.Src)
	d.drawNavMesh(navMesh)
	d.drawSights(navMesh, sights)
	d.mu.Unlock()
	d.event.Send(paint.Event{})
}

func (d *Debugger) close() {
	d.mu.Lock()
	if d.screen != nil {
		d.screen.Release()
		d.screen = nil
	}
	d.mu.Unlock()
}

func (d *Debugger) setScreen(s screen.Buffer) {
	d.mu.Lock()
	if d.screen != nil {
		d.screen.Release()
	}
	d.screen = s
	d.mu.Unlock()
}

func (d *Debugger) uploadScreen(uploader screen.Uploader) {
	d.mu.Lock()
	uploader.Upload(image.ZP, d.screen, d.screen.Bounds())
	d.mu.Unlock()
}

func (d *Debugger) drawNavMesh(navMesh *navmesh.NavMesh) {
	if navMesh == nil {
		d.navMeshView = nil
		return
	}
	if d.navMeshView == nil || d.navMeshView.navMesh != navMesh {
		d.navMeshView = newNavMeshView(navMesh)
	}
	d.navMeshView.draw(d.screen.RGBA())
}

func (d *Debugger) drawSights(navMesh *navmesh.NavMesh, sights []*sight.Sight) {
	if navMesh == nil {
		d.sightViews = d.sightViews[:0]
		return
	}
	for i, s := range sights {
		var sv *sightView
		if len(d.sightViews) > i {
			sv = d.sightViews[i]
			if sv.sight != s {
				sv = newSightView(s, navMesh.Mesh)
				d.sightViews[i] = sv
			}
		} else {
			sv = newSightView(s, navMesh.Mesh)
			d.sightViews = append(d.sightViews, sv)
		}
		sv.draw(d.screen.RGBA())
	}
	d.sightViews = d.sightViews[:len(sights)]
}
