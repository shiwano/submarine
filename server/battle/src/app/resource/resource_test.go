package resource

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestResource(t *testing.T) {
	Convey("resource", t, func() {
		Convey(".clientAssetDir", func() {
			Convey("should exists", func() {
				_, err := os.Stat(clientAssetDir)
				So(err, ShouldBeNil)
			})
		})
	})
}
