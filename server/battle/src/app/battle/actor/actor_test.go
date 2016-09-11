package actor

import (
	"app/battle/context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActor(t *testing.T) {
	Convey("actor", t, func() {
		Convey("should implement context.Actor interface", func() {
			So(&actor{}, ShouldImplement, (*context.Actor)(nil))
		})
	})
}
