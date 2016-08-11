package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

func TestUtil(t *testing.T) {
	Convey("util", t, func() {
		Convey(".EqualFloats", func() {
			quickConfig := &quick.Config{
				MaxCount: 100,
				Rand:     rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
				Values: func(args []reflect.Value, rand *rand.Rand) {
					args[0] = reflect.ValueOf(rand.Float64())
					args[1] = reflect.ValueOf(rand.Float64() + rand.Float64())
				},
			}

			Convey("with same values", func() {
				Convey("should return true", func() {
					So(quick.Check(func(a, b float64) bool {
						return EqualFloats(a, a)
					}, quickConfig), ShouldBeNil)
				})
			})

			Convey("with different values", func() {
				Convey("should return false", func() {
					So(quick.Check(func(a, b float64) bool {
						return !EqualFloats(a, b)
					}, quickConfig), ShouldBeNil)
				})
			})
		})
	})
}
