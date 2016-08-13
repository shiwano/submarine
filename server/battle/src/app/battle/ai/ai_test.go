package ai

import (
	"app/battle/context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAI(t *testing.T) {
	Convey("ai", t, func() {
		Convey("should implement context.AI interface", func() {
			So(&ai{}, ShouldImplement, (*context.AI)(nil))
		})
	})
}
