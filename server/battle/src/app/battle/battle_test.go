package battle

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBattle(t *testing.T) {
	Convey("Battle", t, func() {
		battle := New(60 * time.Second)

		Convey("#CreateSubmarineUnlessExists", func() {
			Convey("when the submarine does not exist", func() {
				Convey("should not create the submarine", func() {
					battle.CreateSubmarineUnlessExists(1)
					submarine := battle.context.container.getSubmarineByUserID(1)
					So(submarine, ShouldNotBeNil)
				})
			})

			Convey("when the submarine already exists", func() {
				battle.CreateSubmarineUnlessExists(1)
				submarine := battle.context.container.getSubmarineByUserID(1)

				Convey("should not replace the existing with new submarine instance", func() {
					battle.CreateSubmarineUnlessExists(1)
					So(submarine, ShouldEqual, battle.context.container.getSubmarineByUserID(1))
				})
			})
		})

		Convey("#update", func() {
			battle.start()

			Convey("when the battle is not finished", func() {
				Convey("should return true", func() {
					isFinished := battle.update(time.Now())
					So(isFinished, ShouldBeTrue)
				})
			})

			Convey("when the battle is finished", func() {
				Convey("should return false", func() {
					isFinished := battle.update(time.Now().Add(60 * time.Second))
					So(isFinished, ShouldBeFalse)
				})
			})
		})
	})
}
