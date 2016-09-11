package navmesh

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestObject(t *testing.T) {
	Convey("object", t, func() {
		mesh, _ := LoadMeshFromJSONFile("fixtures/mesh.json")
		navMesh := New(mesh)
		object := &object{
			id:         1,
			navMesh:    navMesh,
			position:   &vec2.T{1, 1},
			sizeRadius: 3,
		}

		Convey("should implement Object interface", func() {
			So(object, ShouldImplement, (*Object)(nil))
		})

		Convey("#IntersectWithLine", func() {
			Convey("with an intersected points", func() {
				Convey("should return the intersection point", func() {
					p, ok := object.IntersectWithLineSeg(
						&vec2.T{5, 1},
						(&vec2.T{0, 1}).Sub(&vec2.T{5, 1}).Normalize(),
						(&vec2.T{0, 1}).Sub(&vec2.T{5, 1}),
					)
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, 4)
					So(p[1], ShouldEqual, 1)
				})
			})

			Convey("with a no-intersected points", func() {
				Convey("should return nil", func() {
					_, ok := object.IntersectWithLineSeg(
						&vec2.T{5, 1},
						(&vec2.T{7, 1}).Sub(&vec2.T{5, 1}).Normalize(),
						(&vec2.T{7, 1}).Sub(&vec2.T{5, 1}),
					)
					So(ok, ShouldBeFalse)

					_, ok = object.IntersectWithLineSeg(
						&vec2.T{1, 1},
						(&vec2.T{7, 1}).Sub(&vec2.T{1, 1}).Normalize(),
						(&vec2.T{7, 1}).Sub(&vec2.T{1, 1}),
					)
					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
