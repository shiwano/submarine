package component

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimer(t *testing.T) {
	Convey("Timer", t, func() {
		startTime := time.Now()
		t := NewTimer(startTime)

		Convey("#Register", func() {
			Convey("should register timer with the specified interval", func() {
				callCount := 0
				t.Register(1, func() {
					callCount++
				})
				t.Update(startTime.Add(1 * time.Second))
				So(callCount, ShouldEqual, 0)
				t.Update(startTime.Add(1*time.Second + 1))
				So(callCount, ShouldEqual, 1)
				t.Update(startTime.Add(2*time.Second + 1))
				So(callCount, ShouldEqual, 1)
			})

			Convey("should return a timer item that is cancelable", func() {
				callCount := 0
				item := t.Register(1, func() {
					callCount++
				})
				item.Cancel()
				t.Update(startTime.Add(1*time.Second + 1))
				So(callCount, ShouldEqual, 0)
			})
		})

		Convey("#RegisterRepeat", func() {
			Convey("should register repeat timer with the specified interval", func() {
				callCount := 0
				t.RegisterRepeat(1, func() {
					callCount++
				})
				t.Update(startTime.Add(1 * time.Second))
				So(callCount, ShouldEqual, 0)
				t.Update(startTime.Add(1*time.Second + 1))
				So(callCount, ShouldEqual, 1)
				t.Update(startTime.Add(2*time.Second + 1))
				So(callCount, ShouldEqual, 2)
			})

			Convey("should return a timer item that is cancelable", func() {
				callCount := 0
				item := t.RegisterRepeat(1, func() {
					callCount++
				})
				item.Cancel()
				t.Update(startTime.Add(1*time.Second + 1))
				So(callCount, ShouldEqual, 0)
			})
		})
	})
}
