package battle

import (
	"testing"
	"time"

	"github.com/shiwano/submarine/server/battle/src/resource"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBattle(t *testing.T) {
	Convey("Battle", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		lightMap, _ := resource.Loader.LoadLightMap(1)
		b := New(60*time.Second, stageMesh, lightMap)

		Convey("#EnterUser", func() {
			Convey("should create the submarine", func() {
				b.EnterUser(1)
				s, ok := b.scene.SubmarineByPlayerID(1)
				So(ok, ShouldBeTrue)
				So(s, ShouldNotBeNil)
			})

			Convey("when the submarine already exists", func() {
				b.EnterUser(1)

				Convey("should not replace the existing with new submarine instance", func() {
					s, _ := b.scene.SubmarineByPlayerID(1)
					b.EnterUser(1)
					s2, ok := b.scene.SubmarineByPlayerID(1)
					So(ok, ShouldBeTrue)
					So(s, ShouldEqual, s2)
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
