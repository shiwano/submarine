package ai

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/src/battle/scene"
)

func TestAI(t *testing.T) {
	Convey("ai", t, func() {
		Convey("should implement scene.AI interface", func() {
			So(&ai{}, ShouldImplement, (*scene.AI)(nil))
		})
	})
}
