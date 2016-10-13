package resource

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
