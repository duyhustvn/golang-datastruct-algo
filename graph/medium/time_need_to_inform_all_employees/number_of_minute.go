package time_need_to_inform_all_employees

import (
	"ds/stack"
	"fmt"
)

// https://leetcode.com/problems/time-needed-to-inform-all-employees/

type Node struct {
	Value int
}

type ItemGraph struct {
	node  []*Node
	Edges map[Node][]*Node
}

func NewNode(v int) *Node {
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

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjacencyList := buildAdjacencyList(headID, manager)
	grap, headNode := buildGraphFromAdjacencyList(adjacencyList, headID)
	return getDepth(grap, headNode, informTime)
}

func buildAdjacencyList(headID int, manager []int) map[int][]int {
	mapDirectManagerToSub := make(map[int][]int) // map direct manager to subordination

	for suborId, man := range manager {
		if man == -1 {
			continue
		}
		subors, ok := mapDirectManagerToSub[man]
		if !ok {
			mapDirectManagerToSub[man] = []int{suborId}
			continue
		}
		mapDirectManagerToSub[man] = append(subors, suborId)
	}

	return mapDirectManagerToSub
}

func buildGraphFromAdjacencyList(adjacencyList map[int][]int, headID int) (*ItemGraph, *Node) {
	graph := NewGraph()
	var headNode *Node
	existedNode := make(map[int]*Node)
	for man, subors := range adjacencyList {
		manNode, ok := existedNode[man]
		if !ok {
			manNode = NewNode(man)
		}
		if man == headID {
			headNode = manNode
		}
		graph.AddNode(manNode)
		for _, sub := range subors {
			subNode := NewNode(sub)
			graph.AddNode(subNode)
			graph.AddEdge(manNode, subNode)
		}
	}
	return graph, headNode
}

func getDepth(graph *ItemGraph, startNode *Node, informTime []int) int {
	depth := make(map[*Node]int)
	maxDepth := 0
	visited := make(map[Node]bool)
	s := stack.Stack{}
	s.Push(startNode)
	visited[*startNode] = true
	depth[startNode] = 0
	for !s.IsEmpty() {
		node := s.Pop().(*Node)
		vNode := (*node).Value
		neighbors := graph.Edges[*node]
		for _, neighbor := range neighbors {
			if _, ok := visited[*neighbor]; !ok {
				preDepth := depth[node]
				if preDepth+informTime[vNode] > maxDepth {
					maxDepth = preDepth + informTime[vNode]
				}
				depth[neighbor] = preDepth + informTime[vNode]
				s.Push(neighbor)
				visited[*neighbor] = true
			}
		}
	}
	return maxDepth
}
