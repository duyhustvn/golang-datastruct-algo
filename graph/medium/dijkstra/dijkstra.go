package findpath

import (
	priorityqueue "ds/priority_queue"
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
	mutex sync.Mutex
}

func NewGrapth() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func newNode(name string) *Node {
	return &Node{name: name, value: math.MaxInt, through: nil}
}

func newEdge(node *Node, weight int) *Edge {
	return &Edge{node: node, weight: weight}
}

func (g *WeightedGraph) AddNode(node *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Nodes = append(g.Nodes, node)
}

func (g *WeightedGraph) AddNodes(names []string) map[string]*Node {
	nodes := make(map[string]*Node, len(names))
	for _, name := range names {
		node := newNode(name)
		g.AddNode(node)
		nodes[name] = node
	}
	return nodes
}

func (g *WeightedGraph) AddEdge(node1 *Node, node2 *Node, weight int) {
	destEdge := newEdge(node2, weight)
	g.Edges[node1.name] = append(g.Edges[node1.name], destEdge)

	destEdge = newEdge(node1, weight)
	g.Edges[node2.name] = append(g.Edges[node2.name], destEdge)
}

func (g *WeightedGraph) GetNode(city string) *Node {
	for _, node := range g.Nodes {
		if node.name == city {
			return node
		}
	}
	return nil
}

func Dijkstra(g *WeightedGraph, fromCity string, toCity string) {
	queue := priorityqueue.NewPriorityQueue(true)

	startNode := g.GetNode(fromCity)

	queue.Enqueue(startNode, 0)

	visited := make(map[string]bool)

	for !queue.IsEmpty() {
		curr, _ := queue.Dequeue()
		currCity := curr.Value.(*Node).name
		if _, ok := visited[currCity]; !ok {
			visited[currCity] = true

			if currCity == toCity {
				return
			}

			// path = append(path, curr.Value.(string))

			neighbors := g.Edges[currCity]

			for _, neighbor := range neighbors {
				if _, ok := visited[neighbor.node.name]; !ok {
					distance := curr.Priority + neighbor.weight
					queue.Enqueue(neighbor.node, distance)

					if distance < neighbor.node.value {
						neighbor.node.value = distance
						neighbor.node.through = curr.Value.(*Node)
					}
				}
			}
		}
	}

	return
}
