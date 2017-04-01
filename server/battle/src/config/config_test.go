package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	configAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/configuration"
)

func TestConfig(t *testing.T) {
	Convey("config", t, func() {
		Convey(".Config", func() {
			Convey("should be the server config", func() {
				So(Config, ShouldHaveSameTypeAs, &configAPI.Server{})
				So(Config.ApiServerBaseUri, ShouldEqual, "http://localhost:3000")
			})
		})
	})
}
