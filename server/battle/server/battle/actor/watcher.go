package actor

import (
	"github.com/shiwano/submarine/server/battle/lib/navmesh"
	"github.com/shiwano/submarine/server/battle/server/battle/actor/component"
	"github.com/shiwano/submarine/server/battle/server/battle/context"

	"github.com/ungerik/go3d/float64/vec2"
)

type watcher struct {
	*actor
	timer *component.Timer
}

// NewWatcher creates a watcher.
func NewWatcher(ctx context.Context, user *context.Player, position *vec2.T, direction float64) context.Actor {
	w := &watcher{
		actor: newActor(ctx, user, user.WatcherParams, position, direction),
	}
	w.ignoredLayer = navmesh.LayerAll
	w.timer = component.NewTimer(ctx.Now())
	w.timer.Register(user.WatcherParams.UptimeSeconds, w.onElapsedUptime)

	w.ctx.Event().EmitActorCreateEvent(w)
	return w
}

func (w *watcher) Update() {
	w.timer.Update(w.ctx.Now())
}

func (w *watcher) onElapsedUptime() {
	w.Destroy()
}
