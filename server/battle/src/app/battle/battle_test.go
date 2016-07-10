package battle

import (
	"app/resource"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBattle(t *testing.T) {
	Convey("Battle", t, func() {
		stageMesh, _ := resource.Loader.LoadStageMesh(1)
		battle := New(60*time.Second, stageMesh)

		Convey("#EnterUser", func() {
			Convey("should create the submarine", func() {
				battle.EnterUser(1)
				submarine := battle.context.SubmarineByUserID(1)
				So(submarine, ShouldNotBeNil)
			})

			Convey("when the submarine already exists", func() {
				battle.EnterUser(1)
				submarine := battle.context.SubmarineByUserID(1)

				Convey("should not replace the existing with new submarine instance", func() {
					battle.EnterUser(1)
					So(submarine, ShouldEqual, battle.context.SubmarineByUserID(1))
				})
			})

			Convey("when the battle already is running", func() {
				battle.isStarted = true
				battle.start()

				Convey("should send to reenterUserCh", func() {
					battle.EnterUser(1)
					So(len(battle.reenterUserCh), ShouldEqual, 1)
				})
			})
		})
	})
}
