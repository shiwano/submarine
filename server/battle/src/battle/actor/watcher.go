package actor

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/src/battle/actor/component"
	"github.com/shiwano/submarine/server/battle/src/battle/scene"

	"github.com/ungerik/go3d/float64/vec2"
)

type watcher struct {
	*actor
	timer *component.Timer
}

func newWatcher(scn scene.Scene, user *scene.Player, position *vec2.T, direction float64) scene.Actor {
	w := &watcher{
		actor: newActor(scn, user, user.WatcherParams, position, direction),
	}
	w.ignoredLayer = navmesh.LayerAll
	w.timer = component.NewTimer(scn.Now())
	w.timer.Register(user.WatcherParams.UptimeSeconds, w.onElapsedUptime)

	w.scene.Event().EmitActorCreateEvent(w)
	return w
}

func (w *watcher) Update() {
	w.timer.Update(w.scene.Now())
}

func (w *watcher) onElapsedUptime() {
	w.Destroy()
}
