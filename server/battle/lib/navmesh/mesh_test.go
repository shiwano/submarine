package navmesh

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestMesh(t *testing.T) {
	Convey("mesh", t, func() {
		m, _ := LoadMeshFromJSONFile("fixtures/mesh.json")

		Convey(".LoadMeshFromJSONFile", func() {
			Convey("with a valid json path", func() {
				Convey("should creates a mesh", func() {
					So(m.Version, ShouldEqual, "2c51e198ca05edc8c7ef18f2a1a8174c864980d7328e439e2263d1595acadb35")
					So(m.Rect, ShouldResemble, &vec2.Rect{
						Min: vec2.T{-10, -14},
						Max: vec2.T{10, 7},
					})
					So(m.vertices[0], ShouldResemble, &vec2.T{0, 0})
					So(m.Triangles[0], ShouldResemble, newTriangle(
						m.vertices[0],
						m.vertices[1],
						m.vertices[2],
					))
				})
			})

			Convey("should initialize outerEdges", func() {
				So(m.outerEdges, ShouldHaveLength, 8)
				var edgePointsList [][2]*vec2.T
				for _, e := range m.outerEdges {
					edgePointsList = append(edgePointsList, e.points)
				}
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[0], m.vertices[1]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[1], m.vertices[2]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[2], m.vertices[3]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[3], m.vertices[4]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[0], m.vertices[5]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[4], m.vertices[7]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[5], m.vertices[6]})
				So(edgePointsList, ShouldContain, [2]*vec2.T{m.vertices[6], m.vertices[7]})
			})

			Convey("should initialize trianglesByVertex", func() {
				triangles := m.trianglesByVertex[m.vertices[0]]
				So(triangles, ShouldHaveLength, 4)
				So(triangles, ShouldContain, m.Triangles[0])
				So(triangles, ShouldContain, m.Triangles[1])
				So(triangles, ShouldContain, m.Triangles[2])
				So(triangles, ShouldContain, m.Triangles[3])
			})

			Convey("should initialize adjoiningVertices", func() {
				vertices := m.adjoiningVertices[m.vertices[0]]
				So(vertices, ShouldHaveLength, 5)
				So(vertices, ShouldContain, m.vertices[1])
				So(vertices, ShouldContain, m.vertices[2])
				So(vertices, ShouldContain, m.vertices[3])
				So(vertices, ShouldContain, m.vertices[4])
				So(vertices, ShouldContain, m.vertices[5])
			})
		})

		Convey("#findTriangleByPoint", func() {
			Convey("should find the triangle that contains the specified point", func() {
				triangle := m.findTriangleByPoint(&vec2.T{1, -11})
				So(triangle, ShouldEqual, m.Triangles[5])
			})
		})

		Convey("#isIntersectWithLineSeg", func() {
			Convey("with intersected points", func() {
				Convey("should return true", func() {
					p1 := &vec2.T{1, 0}
					p2 := &vec2.T{1, 10}
					So(m.isIntersectedWithLineSeg(p1, p2.Sub(p1)), ShouldBeTrue)
				})
			})

			Convey("with no-intersected points", func() {
				Convey("should return false", func() {
					p1 := &vec2.T{999, 999}
					p2 := &vec2.T{1000, 1000}
					So(m.isIntersectedWithLineSeg(p1, p2.Sub(p1)), ShouldBeFalse)
				})
			})
		})

		Convey("#intersectWithLineSeg", func() {
			Convey("with intersected points", func() {
				Convey("should return the intersection point", func() {
					p1 := &vec2.T{1, 0}
					p2 := &vec2.T{10, 0}
					p, ok := m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, 8.875)
					So(p[1], ShouldEqual, 0)
				})
			})

			Convey("with intersected points which has multi intersection points", func() {
				Convey("should return the most neaby intersection point", func() {
					p1 := &vec2.T{-7, 3.5}
					p2 := &vec2.T{-1, -11}
					p, ok := m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, -5.551724137931035)
					So(p[1], ShouldEqual, 0)

					p1 = &vec2.T{0, 0}
					p2 = &vec2.T{0, -100}
					p, ok = m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeTrue)
					So(p[0], ShouldEqual, 0)
					So(p[1], ShouldEqual, -14)
				})
			})

			Convey("with no-intersected points", func() {
				Convey("should return nil", func() {
					p1 := &vec2.T{0.00000000001, 0}
					p2 := &vec2.T{5, 0}
					_, ok := m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)

					p1 = &vec2.T{3, 2}
					p2 = &vec2.T{2, -10}
					_, ok = m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)

					p1 = &vec2.T{0, 0}
					p2 = &vec2.T{0, -3}
					_, ok = m.intersectWithLineSeg(p1, p2.Sub(p1))
					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
