package lib

import (
	"ds/queue"
	"ds/stack"
	"fmt"
)

type Vertex struct {
	Key       int
	Verticles map[int]*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{Key: key, Verticles: map[int]*Vertex{}}
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewDirectedGraph() *Graph {
	return &Graph{Vertices: map[int]*Vertex{}, directed: true}
}

func NewUndirectedGraph() *Graph {
	return &Graph{Vertices: map[int]*Vertex{}, directed: false}
}

func (g *Graph) AddVertex(key int) {
	vertex := NewVertex(key)
	g.Vertices[key] = vertex
}

func (g *Graph) AddEdge(k1, k2 int) {
	// get vetex from graph
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		panic("vertices not exists")
	}

	// if v1 already connected to v2
	// do nothing
	if _, ok := v1.Verticles[v2.Key]; ok {
		return
	}

	v1.Verticles[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Verticles[v1.Key] = v1
	}
}

type Node struct {
	Value interface{}
}

type ItemGraph struct {
	node  []*Node
	Edges map[Node][]*Node
}

func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

func NewGraph() *ItemGraph {
	return &ItemGraph{node: make([]*Node, 0), Edges: make(map[Node][]*Node)}
}

func (g *ItemGraph) AddNode(node *Node) {
	g.node = append(g.node, node)
}

func (g *ItemGraph) AddEdge(n1 *Node, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	g.Edges[*n1] = append(g.Edges[*n1], n2)
	g.Edges[*n2] = append(g.Edges[*n2], n1)
}

func (g *ItemGraph) String() {
	for _, node := range g.node {
		vertices := g.Edges[*node]
		fmt.Print(node.Value, " -> ")
		for _, vetex := range vertices {
			fmt.Print(vetex.Value, " ")
		}
		fmt.Println()
	}
}

func BFS(g *ItemGraph, startNode *Node) []interface{} {
	result := make([]interface{}, 0)
	q := queue.Queue{}
	visited := make(map[Node]bool)

	q.Enqueue(startNode)
	visited[*startNode] = true

	for !q.IsEmpty() {
		nodeInQueue, _ := q.Dequeue()
		result = append(result, nodeInQueue.(*Node).Value)
		verticles := g.Edges[*nodeInQueue.(*Node)]

		for _, vertex := range verticles {
			if _, ok := visited[*vertex]; !ok {
				visited[*vertex] = true
				q.Enqueue(vertex)
			}
		}

	}
	return result
}

func DFS(g *ItemGraph, startNode *Node) []interface{} {
	result := make([]interface{}, 0)
	visited := make(map[Node]bool)

	s := stack.Stack{}
	s.Push(startNode)
	visited[*startNode] = true

	for !s.IsEmpty() {
		vertex := s.Pop()
		result = append(result, vertex.(*Node).Value)
		neighbors := g.Edges[*vertex.(*Node)]
		for _, neighbor := range neighbors {
			if _, ok := visited[*neighbor]; !ok {
				visited[*neighbor] = true
				s.Push(neighbor)
			}
		}
	}

	return result
}
