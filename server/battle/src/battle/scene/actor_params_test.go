package scene

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActorParams(t *testing.T) {
	Convey("actorParams", t, func() {
		Convey("should implement ActorParams interface", func() {
			So(&actorParams{}, ShouldImplement, (*ActorParams)(nil))
		})
	})
}
