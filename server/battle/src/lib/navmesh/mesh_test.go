package navmesh

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"testing"
)

func TestMesh(t *testing.T) {
	Convey("mesh", t, func() {
		m, _ := LoadMeshFromJSONFile("fixtures/mesh.json")

		Convey(".LoadMeshFromJSONFile", func() {
			Convey("with a valid json path", func() {
				Convey("should creates a mesh", func() {
					So(m.Vertices[0], ShouldResemble, &vec2.T{0, 0})
					So(m.Triangles[0], ShouldResemble, newTriangle(
						m.Vertices[0],
						m.Vertices[1],
						m.Vertices[2],
					))
				})
			})

			Convey("should initialize outerEdges", func() {
				So(m.outerEdges, ShouldHaveLength, 8)
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[0], m.Vertices[1]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[1], m.Vertices[2]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[2], m.Vertices[3]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[3], m.Vertices[4]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[0], m.Vertices[5]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[4], m.Vertices[7]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[5], m.Vertices[6]})
				So(m.outerEdges, ShouldContain, Edge{m.Vertices[6], m.Vertices[7]})
			})

			Convey("should initialize trianglesByVertex", func() {
				triangles := m.trianglesByVertex[m.Vertices[0]]
				So(triangles, ShouldHaveLength, 4)
				So(triangles, ShouldContain, m.Triangles[0])
				So(triangles, ShouldContain, m.Triangles[1])
				So(triangles, ShouldContain, m.Triangles[2])
				So(triangles, ShouldContain, m.Triangles[3])
			})

			Convey("should initialize adjoiningVertices", func() {
				vertices := m.adjoiningVertices[m.Vertices[0]]
				So(vertices, ShouldHaveLength, 5)
				So(vertices, ShouldContain, m.Vertices[1])
				So(vertices, ShouldContain, m.Vertices[2])
				So(vertices, ShouldContain, m.Vertices[3])
				So(vertices, ShouldContain, m.Vertices[4])
				So(vertices, ShouldContain, m.Vertices[5])
			})
		})

		Convey("#findTriangleByPoint", func() {
			Convey("should find the triangle that contains the specified point", func() {
				triangle := m.findTriangleByPoint(&vec2.T{1, -11})
				So(triangle, ShouldEqual, m.Triangles[5])
			})
		})

		Convey("#intersect", func() {
			Convey("with intersected points", func() {
				Convey("should return the intersection point", func() {
					p1 := &vec2.T{0, 0}
					p2 := &vec2.T{10, 0}
					So(m.intersect(p1, p2), ShouldNotBeNil)
				})
			})

			Convey("with no-intersected points", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{0.00000000001, 0}
					p2 := &vec2.T{5, 0}
					So(m.intersect(p1, p2), ShouldBeNil)

					p1 = &vec2.T{3, 2}
					p2 = &vec2.T{2, -10}
					So(m.intersect(p1, p2), ShouldBeNil)
				})
			})
		})
	})
}
