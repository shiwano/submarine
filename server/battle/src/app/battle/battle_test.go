package battle

import (
	"app/resource"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBattle(t *testing.T) {
	Convey("Battle", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		b := New(60*time.Second, stageMesh)

		Convey("#EnterUser", func() {
			Convey("should create the submarine", func() {
				b.EnterUser(1)
				submarine := b.ctx.SubmarineByPlayerID(1)
				So(submarine, ShouldNotBeNil)
			})

			Convey("when the submarine already exists", func() {
				b.EnterUser(1)
				submarine := b.ctx.SubmarineByPlayerID(1)

				Convey("should not replace the existing with new submarine instance", func() {
					b.EnterUser(1)
					So(submarine, ShouldEqual, b.ctx.SubmarineByPlayerID(1))
				})
			})

			Convey("when the battle already is running", func() {
				b.isStarted = true
				b.start()

				Convey("should send to reenterUserCh", func() {
					b.EnterUser(1)
					So(len(b.reenterUserCh), ShouldEqual, 1)
				})
			})
		})
	})
}
