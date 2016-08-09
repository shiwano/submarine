package actor

import (
	"app/battle/context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestActor(t *testing.T) {
	Convey("actor", t, func() {
		Convey("should implement context.Actor interface", func() {
			So(&actor{}, ShouldImplement, (*context.Actor)(nil))
		})
	})
}
