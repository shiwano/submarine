package resource

import (
	"github.com/k0kubun/pp"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

var p = pp.Println

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
