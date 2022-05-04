package graph

import (
	"testing"
)

func TestAddEdge(t *testing.T) {
	// g := NewDirectedGraph()
	// g.AddVertex(1)
	// g.AddVertex(2)
	// g.AddVertex(3)
	// g.AddVertex(4)

	// g.AddEdge(1, 2)
	// g.AddEdge(2, 3)
	// g.AddEdge(2, 4)
	// g.AddEdge(4, 1)

	// t.Logf("vertices: %+v\n", g.Vertices)
	// assert.NotNil(t, g.Vertices)

	n1 := NewNode(1)
	n5 := NewNode(5)
	n4 := NewNode(4)
	n2 := NewNode(2)
	n7 := NewNode(7)
	n6 := NewNode(6)
	n3 := NewNode(3)
	n8 := NewNode(8)

	g := NewGraph()

	g.AddNode(n1)
	g.AddNode(n5)
	g.AddNode(n4)
	g.AddNode(n2)
	g.AddNode(n7)
	g.AddNode(n6)
	g.AddNode(n3)
	g.AddNode(n8)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n4)
	g.AddEdge(n1, n5)
	g.AddEdge(n2, n7)
	g.AddEdge(n2, n3)
	g.AddEdge(n3, n8)
	g.AddEdge(n2, n6)

	g.String()

	result := g.BFS()
	t.Logf("result: %+v", result)
}
