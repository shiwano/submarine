package actor

import (
	"testing"

	"github.com/shiwano/submarine/server/battle/src/battle/context"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActor(t *testing.T) {
	Convey("actor", t, func() {
		Convey("should implement context.Actor interface", func() {
			So(&actor{}, ShouldImplement, (*context.Actor)(nil))
		})
	})
}
