package main

import (
	"fmt"
	"math"
	"sync"
)

type Node struct {
	name    string
	value   int
	through *Node
}

type Edge struct {
	node   *Node
	weight int
}

type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *WeightedGraph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

func (g *WeightedGraph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(graph *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight})
}

func buildGraph() *WeightedGraph {
	graph := NewGraph()
	nodes := AddNodes(graph,
		"London",
		"Paris",
		"Amsterdam",
		"Luxembourg",
		"Zurich",
		"Rome",
		"Berlin",
		"Vienna",
		"Warsaw",
		"Istanbul",
	)
	graph.AddEdge(nodes["London"], nodes["Paris"], 80)
	graph.AddEdge(nodes["London"], nodes["Luxembourg"], 75)
	graph.AddEdge(nodes["London"], nodes["Amsterdam"], 75)
	graph.AddEdge(nodes["Paris"], nodes["Luxembourg"], 60)
	graph.AddEdge(nodes["Paris"], nodes["Rome"], 125)
	graph.AddEdge(nodes["Luxembourg"], nodes["Berlin"], 90)
	graph.AddEdge(nodes["Luxembourg"], nodes["Zurich"], 60)
	graph.AddEdge(nodes["Luxembourg"], nodes["Amsterdam"], 55)
	graph.AddEdge(nodes["Zurich"], nodes["Vienna"], 80)
	graph.AddEdge(nodes["Zurich"], nodes["Rome"], 90)
	graph.AddEdge(nodes["Zurich"], nodes["Berlin"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Amsterdam"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Vienna"], 75)
	graph.AddEdge(nodes["Vienna"], nodes["Rome"], 100)
	graph.AddEdge(nodes["Vienna"], nodes["Istanbul"], 130)
	graph.AddEdge(nodes["Warsaw"], nodes["Berlin"], 80)
	graph.AddEdge(nodes["Warsaw"], nodes["Istanbul"], 180)
	graph.AddEdge(nodes["Rome"], nodes["Istanbul"], 155)

	return graph
}

func main() {
	g := buildGraph()
	for _, node := range g.Nodes {
		fmt.Printf("node: %+v", node)
	}
	fmt.Println()
	for k, v := range g.Edges {
		fmt.Printf("k: %+v v: %+v\n", k, v)
		for _, e := range v {
			fmt.Printf("e: %+v node: %+v\n", e, e.node)
		}
	}
}
