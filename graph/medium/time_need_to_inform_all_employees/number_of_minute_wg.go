package time_need_to_inform_all_employees

import "fmt"

// Weighted Graph
type GNode struct {
	Value int
}

type GEdge struct {
	Node   *GNode
	Weight int
}

type WeightedGraph struct {
	Nodes []*GNode
	Edges map[GNode][]*GEdge
}

func (g *WeightedGraph) String() {
	for _, node := range g.Nodes {
		edges := g.Edges[*node]
		fmt.Print(node.Value, " -> ")
		for _, edge := range edges {
			fmt.Print(edge.Node.Value, ":", edge.Weight, " ")
		}
		fmt.Println()
	}
}

func NewGNode(v int) *GNode {
	return &GNode{Value: v}
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{Nodes: []*GNode{}, Edges: make(map[GNode][]*GEdge)}
}

func (g *WeightedGraph) AddNode(n *GNode) {
	g.Nodes = append(g.Nodes, n)
}

func (g *WeightedGraph) AddEdge(n1 *GNode, n2 *GNode, weigh int) {
	g.Edges[*n1] = append(g.Edges[*n1], &GEdge{n2, weigh})
	g.Edges[*n2] = append(g.Edges[*n2], &GEdge{n1, weigh})
}

type AdjacencyListWithWeight struct {
	Sub    []int
	Weight int
}

func buildAdjacencyListWithWeight(managers []int, time2Inform []int) map[int]AdjacencyListWithWeight {
	adjacencyListWithWeight := make(map[int]AdjacencyListWithWeight)

	for subID, man := range managers {
		if man == -1 {
			continue
		}

		if adj, ok := adjacencyListWithWeight[man]; !ok {
			adjacencyListWithWeight[man] = AdjacencyListWithWeight{[]int{subID}, time2Inform[man]}
		} else {
			adj.Sub = append(adjacencyListWithWeight[man].Sub, subID)

			adjacencyListWithWeight[man] = adj
		}
	}

	return adjacencyListWithWeight
}

func buildWeightGraphFromAdjacencyList(headerID int, adjacencyList map[int]AdjacencyListWithWeight) (*WeightedGraph, *GNode) {
	g := NewWeightedGraph()
	var headNode *GNode
	for mangagerID, sub := range adjacencyList {
		managerNode := NewGNode(mangagerID)

		if mangagerID == headerID {
			headNode = managerNode
		}

		g.AddNode(managerNode)
		for _, subID := range sub.Sub {
			subNode := NewGNode(subID)
			g.AddNode(subNode)
			g.AddEdge(managerNode, subNode, sub.Weight)
		}
	}
	return g, headNode
}

func numOfMinutesWeightedGraph(n int, headID int, manager []int, informTime []int) int {
	if n == 0 || n == 1 {
		return 0
	}
	adjacencyList := buildAdjacencyListWithWeight(manager, informTime)
	wg, headerNode := buildWeightGraphFromAdjacencyList(headID, adjacencyList)
	return getDepthWeightedGraph(wg, headerNode)
}

func getDepthWeightedGraph(graph *WeightedGraph, startNode *GNode) int {
	s := Stack{}
	visited := make(map[GNode]bool)
	depth := make(map[GNode]int)
	s.Push(startNode)
	depth[*startNode] = 0
	max := 0

	for !s.IsEmpty() {
		node := s.Pop().(*GNode)
		for _, edge := range graph.Edges[*node] {
			if _, ok := visited[*edge.Node]; !ok {
				depth[*edge.Node] = depth[*node] + edge.Weight
				if max < depth[*edge.Node] {
					max = depth[*edge.Node]
				}
				visited[*node] = true
				s.Push(edge.Node)
			}
		}
	}
	return max
}
