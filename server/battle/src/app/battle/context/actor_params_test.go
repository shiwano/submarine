package context

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestActorParams(t *testing.T) {
	Convey("actorParams", t, func() {
		Convey("should implement ActorParams interface", func() {
			So(&actorParams{}, ShouldImplement, (*ActorParams)(nil))
		})
	})
}
