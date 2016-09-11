package navmesh

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/ungerik/go3d/float64/vec2"
)

type vertexData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type meshData struct {
	Vertices  []vertexData `json:"vertices"`
	Triangles [][3]int     `json:"triangles"`
}

// Mesh represents a navmesh.
type Mesh struct {
	Rect              *vec2.Rect
	vertices          []*vec2.T
	triangles         []*Triangle
	outerEdges        []*edge
	trianglesByVertex map[*vec2.T][]*Triangle
	adjoiningVertices map[*vec2.T][]*vec2.T
	distancesByVertex map[*vec2.T]map[*vec2.T]float64
}

// LoadMeshFromJSONFile loads a mesh data from a JSON file.
func LoadMeshFromJSONFile(jsonPath string) (*Mesh, error) {
	f, err := os.Open(jsonPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	rawMesh := new(meshData)
	if err := json.NewDecoder(f).Decode(rawMesh); err != nil {
		return nil, err
	}

	// Vertices
	verticesLength := len(rawMesh.Vertices)
	m := new(Mesh)
	m.Rect = new(vec2.Rect)
	m.vertices = make([]*vec2.T, verticesLength)
	for i, v := range rawMesh.Vertices {
		m.vertices[i] = &vec2.T{v.X, v.Y}

		if v.X < m.Rect.Min[0] {
			m.Rect.Min[0] = v.X
		}
		if v.Y < m.Rect.Min[1] {
			m.Rect.Min[1] = v.Y
		}
		if v.X > m.Rect.Max[0] {
			m.Rect.Max[0] = v.X
		}
		if v.Y > m.Rect.Max[1] {
			m.Rect.Max[1] = v.Y
		}
	}

	// Triangles
	m.triangles = make([]*Triangle, len(rawMesh.Triangles))
	for i, t := range rawMesh.Triangles {
		if t[0] >= verticesLength || t[1] >= verticesLength || t[2] >= verticesLength {
			return nil, fmt.Errorf("Invalid vertex index: %v (%v)", t, i)
		}
		m.triangles[i] = newTriangle(
			m.vertices[t[0]],
			m.vertices[t[1]],
			m.vertices[t[2]],
		)
	}

	// OuterEdges
	rawEdges := make(map[[2]int]int)
	for _, t := range rawMesh.Triangles {
		for _, vertexIndex1 := range t {
			for _, vertexIndex2 := range t {
				if vertexIndex1 != vertexIndex2 {
					if vertexIndex1 <= vertexIndex2 {
						rawEdges[[2]int{vertexIndex1, vertexIndex2}]++
					} else {
						rawEdges[[2]int{vertexIndex2, vertexIndex1}]++
					}
				}
			}
		}
	}
	m.outerEdges = make([]*edge, 0)
	for key, value := range rawEdges {
		if value == 2 {
			a := m.vertices[key[0]]
			b := m.vertices[key[1]]
			for _, triangle := range m.triangles {
				if aIndex, ok := triangle.vertexIndex(a); ok {
					if bIndex, ok := triangle.vertexIndex(b); ok {
						m.outerEdges = append(m.outerEdges, newEdge(triangle, aIndex, bIndex))
					}
				}
			}
		}
	}

	// TrianglesByVertex
	m.trianglesByVertex = make(map[*vec2.T][]*Triangle)
	for _, v := range m.vertices {
		m.trianglesByVertex[v] = m.findTrianglesByVertex(v)
	}

	// AdjoiningVertices
	m.adjoiningVertices = make(map[*vec2.T][]*vec2.T)
	for _, v := range m.vertices {
		m.adjoiningVertices[v] = m.findAdjoiningVertices(v)
	}

	// DistancesByVertex
	m.distancesByVertex = make(map[*vec2.T]map[*vec2.T]float64)
	for _, v1 := range m.vertices {
		m.distancesByVertex[v1] = make(map[*vec2.T]float64)
		for _, v2 := range m.vertices {
			if v1 != v2 {
				diff := vec2.Sub(v1, v2)
				m.distancesByVertex[v1][v2] = diff.Length()
			}
		}
	}
	return m, nil
}

func (m *Mesh) findTrianglesByVertex(vertex *vec2.T) []*Triangle {
	var triangles []*Triangle
	for _, t := range m.triangles {
		if vertex == t.Vertices[0] ||
			vertex == t.Vertices[1] ||
			vertex == t.Vertices[2] {
			triangles = append(triangles, t)
		}
	}
	return triangles
}

func (m *Mesh) findAdjoiningVertices(vertex *vec2.T) []*vec2.T {
	triangles := m.trianglesByVertex[vertex]
	vertexSet := make(map[*vec2.T]struct{})
	for _, t := range triangles {
		for _, v := range t.Vertices {
			if vertex != v {
				vertexSet[v] = struct{}{}
			}
		}
	}

	i, vertices := 0, make([]*vec2.T, len(vertexSet))
	for key := range vertexSet {
		vertices[i] = key
		i++
	}
	return vertices
}

func (m *Mesh) findTriangleByPoint(point *vec2.T) *Triangle {
	for _, t := range m.triangles {
		if t.containsPoint(point) {
			return t
		}
	}
	return nil
}

func (m *Mesh) getOrCalculateDistance(from, to *vec2.T) float64 {
	if filteredDistances, ok := m.distancesByVertex[from]; ok {
		if distance, ok := filteredDistances[to]; ok {
			return distance
		}
	}
	diff := vec2.Sub(to, from)
	return diff.Length()
}

func (m *Mesh) isIntersectedWithLineSeg(lineOrigin, lineVec *vec2.T) bool {
	for _, edge := range m.outerEdges {
		if _, ok := edge.intersectWithLineSeg(lineOrigin, lineVec); ok {
			return true
		}
	}
	return false
}

func (m *Mesh) intersectWithLineSeg(lineOrigin, lineVec *vec2.T) (resultPoint vec2.T, result bool) {
	resultLengthSqr := math.MaxFloat64
	var resultEdgeEndPoint *vec2.T

	for _, edge := range m.outerEdges {
		if p, ok := edge.intersectWithLineSeg(lineOrigin, lineVec); ok {
			lengthSqr := calculateVectorLengthSqr(lineOrigin, &p)

			if resultEdgeEndPoint != nil && equalVectors(resultEdgeEndPoint, &p) {
				resultLengthSqr = lengthSqr
				resultPoint = p
				result = true
				resultEdgeEndPoint = nil
				continue
			}

			if lengthSqr < resultLengthSqr {
				if edge.isEndPoint(&p) {
					resultEdgeEndPoint = &p
					continue
				}

				resultLengthSqr = lengthSqr
				resultPoint = p
				result = true
				resultEdgeEndPoint = nil
			}
		}
	}
	return
}
