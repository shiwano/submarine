package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
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
					result := object.IntersectWithLine(
						&vec2.T{5, 1},
						vec2.Sub(&vec2.T{0, 1}, &vec2.T{5, 1}),
					)
					So(result[0], ShouldEqual, 4)
					So(result[1], ShouldEqual, 1)
				})
			})

			Convey("with an no-intersected points", func() {
				Convey("should return nil", func() {
					result := object.IntersectWithLine(
						&vec2.T{5, 1},
						vec2.Sub(&vec2.T{7, 1}, &vec2.T{5, 1}),
					)
					So(result, ShouldBeNil)

					result = object.IntersectWithLine(
						&vec2.T{1, 1},
						vec2.Sub(&vec2.T{7, 1}, &vec2.T{1, 1}),
					)
					So(result, ShouldBeNil)
				})
			})
		})
	})
}
