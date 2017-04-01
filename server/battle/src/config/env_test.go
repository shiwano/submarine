package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnv(t *testing.T) {
	Convey("config", t, func() {
		Convey(".Env", func() {
			Convey("should be the server env", func() {
				So(Env, ShouldEqual, "test")
			})
		})
	})
}
