package graph

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
	value interface{}
}

type ItemGraph struct {
	node  []*Node
	edges map[Node][]*Node
}

func NewNode(v interface{}) *Node {
	return &Node{value: v}
}

func NewGraph() *ItemGraph {
	return &ItemGraph{node: make([]*Node, 0), edges: make(map[Node][]*Node)}
}

func (g *ItemGraph) AddNode(node *Node) {
	g.node = append(g.node, node)
}

func (g *ItemGraph) AddEdge(n1 *Node, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}

	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *ItemGraph) String() {
	for _, node := range g.node {
		vertices := g.edges[*node]
		fmt.Print(node.value, " -> ")
		for _, vetex := range vertices {
			fmt.Print(vetex.value, " ")
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
		result = append(result, nodeInQueue.(*Node).value)
		verticles := g.edges[*nodeInQueue.(*Node)]

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
		result = append(result, vertex.(*Node).value)
		neighbors := g.edges[*vertex.(*Node)]
		for _, neighbor := range neighbors {
			if _, ok := visited[*neighbor]; !ok {
				visited[*neighbor] = true
				s.Push(neighbor)
			}

		}
	}

	return result
}
