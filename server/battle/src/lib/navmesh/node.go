package navmesh

import (
	"container/heap"

	"github.com/ungerik/go3d/float64/vec2"
)

type node struct {
	parent *node
	point  *vec2.T
	gScore float64
	fScore float64
}

type nodeHeap struct {
	nodes []*node
}

func newNodeHeap(nodes []*node) *nodeHeap {
	n := new(nodeHeap)
	n.nodes = nodes
	heap.Init(n)
	return n
}

func (n *nodeHeap) Len() int {
	return len(n.nodes)
}

func (n *nodeHeap) Less(i, j int) bool {
	return n.nodes[i].fScore < n.nodes[j].fScore
}

func (n *nodeHeap) Swap(i, j int) {
	n.nodes[i], n.nodes[j] = n.nodes[j], n.nodes[i]
}

// Push an element.
func (n *nodeHeap) Push(x interface{}) {
	n.nodes = append(n.nodes, x.(*node))
}

// Pop an element.
func (n *nodeHeap) Pop() interface{} {
	length := len(n.nodes)
	x := n.nodes[length-1]
	n.nodes = n.nodes[:length-1]
	return x
}

func (n *nodeHeap) pushNode(node *node) {
	heap.Push(n, node)
}

func (n *nodeHeap) popNode() *node {
	return heap.Pop(n).(*node)
}
