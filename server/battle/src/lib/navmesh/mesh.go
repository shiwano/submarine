package navmesh

import (
	"encoding/json"
	"fmt"
	"github.com/ungerik/go3d/float64/vec2"
	"os"
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
	Vertices          []*vec2.T
	Triangles         []*Triangle
	outerEdges        []Edge
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
	m.Vertices = make([]*vec2.T, verticesLength)
	for i, v := range rawMesh.Vertices {
		m.Vertices[i] = &vec2.T{v.X, v.Y}
	}

	// Triangles
	m.Triangles = make([]*Triangle, len(rawMesh.Triangles))
	for i, t := range rawMesh.Triangles {
		if t[0] >= verticesLength || t[1] >= verticesLength || t[2] >= verticesLength {
			return nil, fmt.Errorf("Invalid vertex index: %v (%v)", t, i)
		}
		m.Triangles[i] = newTriangle(
			m.Vertices[t[0]],
			m.Vertices[t[1]],
			m.Vertices[t[2]],
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
	m.outerEdges = make([]Edge, 0)
	for key, value := range rawEdges {
		if value == 2 {
			m.outerEdges = append(m.outerEdges, Edge{
				m.Vertices[key[0]],
				m.Vertices[key[1]],
			})
		}
	}

	// TrianglesByVertex
	m.trianglesByVertex = make(map[*vec2.T][]*Triangle)
	for _, v := range m.Vertices {
		m.trianglesByVertex[v] = m.findTrianglesByVertex(v)
	}

	// AdjoiningVertices
	m.adjoiningVertices = make(map[*vec2.T][]*vec2.T)
	for _, v := range m.Vertices {
		m.adjoiningVertices[v] = m.findAdjoiningVertices(v)
	}

	// DistancesByVertex
	m.distancesByVertex = make(map[*vec2.T]map[*vec2.T]float64)
	for _, v1 := range m.Vertices {
		m.distancesByVertex[v1] = make(map[*vec2.T]float64)
		for _, v2 := range m.Vertices {
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
	for _, t := range m.Triangles {
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
	for _, t := range m.Triangles {
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

func (m Mesh) isIntersectedWithLine(lineOrigin *vec2.T, lineVector vec2.T) bool {
	for _, edge := range m.outerEdges {
		if edge.intersectWithLine(lineOrigin, lineVector) != nil {
			return true
		}
	}
	return false
}

func (m Mesh) intersectWithLine(lineOrigin *vec2.T, lineVector vec2.T) *vec2.T {
	var result *vec2.T
	for _, edge := range m.outerEdges {
		if point := edge.intersectWithLine(lineOrigin, lineVector); point != nil {
			if result == nil || point.LengthSqr() < result.LengthSqr() {
				result = point
			}
		}
	}
	return result
}
