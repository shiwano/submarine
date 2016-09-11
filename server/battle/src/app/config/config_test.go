package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("config", t, func() {
		Convey(".Config", func() {
			Convey("should be the server config", func() {
				So(Config, ShouldHaveSameTypeAs, &ServerConfig{})
				So(Config.APIServerBaseURI, ShouldEqual, "http://localhost:3000")
			})
		})
	})
}
