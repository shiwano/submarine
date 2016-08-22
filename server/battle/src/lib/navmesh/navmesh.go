package navmesh

import (
	"github.com/ungerik/go3d/float64/vec2"
)

// RaycastHitInfo represents a raycast hit info.
type RaycastHitInfo struct {
	Object Object
	Point  vec2.T
}

// NavMesh represents a navmesh.
type NavMesh struct {
	Mesh                *Mesh
	HeuristicFunc       func(from, to *vec2.T) float64
	Objects             map[int64]Object
	lastCreatedObjectID int64
}

// New creates a nevmesh.
func New(mesh *Mesh) *NavMesh {
	return &NavMesh{
		Mesh:          mesh,
		HeuristicFunc: calculateOctileDistance,
		Objects:       make(map[int64]Object),
	}
}

// CreateAgent creates an agent.
func (n *NavMesh) CreateAgent(size float64, position *vec2.T) *Agent {
	n.lastCreatedObjectID++
	agent := &Agent{
		object: &object{
			id:         n.lastCreatedObjectID,
			navMesh:    n,
			position:   position,
			sizeRadius: size / 2,
		},
	}
	n.Objects[agent.id] = agent
	return agent
}

// DestroyObject destroys the specified object.
func (n *NavMesh) DestroyObject(objectID int64) {
	delete(n.Objects, objectID)
}

// Raycast casts a ray on the navmesh.
func (n *NavMesh) Raycast(origin, vec *vec2.T, ignoredLayer LayerMask) (hitInfo *RaycastHitInfo, result bool) {
	var resultLengthSqr float64

	if p, ok := n.Mesh.intersectWithLineSeg(origin, vec); ok {
		hitInfo = &RaycastHitInfo{Point: p}
		resultLengthSqr = calculateVectorLengthSqr(origin, &p)
		result = true
	}

	dir := vec.Normalized()
	for _, obj := range n.Objects {
		if !obj.Layer().Has(ignoredLayer) {
			if p, ok := obj.IntersectWithLineSeg(origin, &dir, vec); ok {
				lengthSqr := calculateVectorLengthSqr(origin, &p)

				if !result || lengthSqr < resultLengthSqr {
					hitInfo = &RaycastHitInfo{Object: obj, Point: p}
					resultLengthSqr = lengthSqr
					result = true
				}
			}
		}
	}
	return
}

// FindPath finds a path on the navmesh.
func (n *NavMesh) FindPath(start *vec2.T, goal *vec2.T) []vec2.T {
	startTriangle := n.Mesh.findTriangleByPoint(start)
	goalTriangle := n.Mesh.findTriangleByPoint(goal)

	if startTriangle == nil || goalTriangle == nil {
		return []vec2.T{}
	}

	goalVector := vec2.Sub(goal, start)
	if startTriangle == goalTriangle ||
		!n.Mesh.isIntersectedWithLineSeg(start, &goalVector) {
		return []vec2.T{*start, *goal}
	}

	closedPoints := make(map[*vec2.T]struct{})
	nodeHeap := newNodeHeap([]*node{
		&node{
			parent: nil,
			point:  start,
			gScore: 0,
			fScore: 0,
		},
	})

	for nodeHeap.Len() > 0 {
		current := nodeHeap.popNode()
		if current.point == goal {
			return n.makePathFromNode(current)
		}
		closedPoints[current.point] = struct{}{}

		var nearbyPoints []*vec2.T
		if current.point == start {
			nearbyPoints = []*vec2.T(startTriangle.Vertices[:])
		} else if goalTriangle.hasVertex(current.point) {
			nearbyPoints = []*vec2.T{goal}
		} else {
			nearbyPoints = n.Mesh.adjoiningVertices[current.point]
		}

		for _, nearbyPoint := range nearbyPoints {
			if _, ok := closedPoints[nearbyPoint]; ok {
				continue
			}

			gScore := current.gScore + n.Mesh.getOrCalculateDistance(current.point, nearbyPoint)
			nodeHeap.pushNode(&node{
				parent: current,
				point:  nearbyPoint,
				gScore: gScore,
				fScore: gScore + n.HeuristicFunc(nearbyPoint, goal),
			})
		}
	}
	return []vec2.T{}
}

func (n *NavMesh) makePathFromNode(node *node) []vec2.T {
	var path []vec2.T
	current := node
	for {
		path = append(path, *current.point)
		if current.parent == nil {
			break
		}
		current = current.parent
	}
	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
	return path
}
