package actor

import (
	"testing"

	"github.com/shiwano/submarine/server/battle/src/battle/scene"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActor(t *testing.T) {
	Convey("actor", t, func() {
		Convey("should implement scene.Actor interface", func() {
			So(&actor{}, ShouldImplement, (*scene.Actor)(nil))
		})
	})
}
