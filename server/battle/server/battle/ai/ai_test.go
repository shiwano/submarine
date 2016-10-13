package ai

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/shiwano/submarine/server/battle/server/battle/context"
)

func TestAI(t *testing.T) {
	Convey("ai", t, func() {
		Convey("should implement context.AI interface", func() {
			So(&ai{}, ShouldImplement, (*context.AI)(nil))
		})
	})
}
