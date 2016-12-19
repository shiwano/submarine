package component

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiLock(t *testing.T) {
	Convey("MultiLock", t, func() {
		l := new(MultiLock)

		Convey("#Lock", func() {
			Convey("should lock", func() {
				l.Lock()
				So(l.IsLocked(), ShouldBeTrue)
			})

			Convey("should lock in duplicate", func() {
				l.Lock()
				l.Lock()
				l.Unlock()
				So(l.IsLocked(), ShouldBeTrue)
			})
		})

		Convey("#Unlock", func() {
			Convey("should unlock", func() {
				l.Lock()
				l.Unlock()
				So(l.IsLocked(), ShouldBeFalse)
			})

			Convey("should not unlock in duplicate", func() {
				l.Unlock()
				l.Unlock()
				l.Lock()
				So(l.IsLocked(), ShouldBeTrue)
			})
		})
	})
}
