package ai

import (
	"app/battle/context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAI(t *testing.T) {
	Convey("ai", t, func() {
		Convey("should implement context.AI interface", func() {
			So(&ai{}, ShouldImplement, (*context.AI)(nil))
		})
	})
}
