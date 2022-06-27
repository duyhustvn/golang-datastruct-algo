package findpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNodes(t *testing.T) {
	graph := NewGrapth()

	nodes := graph.AddNodes([]string{
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
	})

	assert.Equal(t, "London", graph.Nodes[0].name)
	assert.Equal(t, "Paris", graph.Nodes[1].name)
	assert.Equal(t, "Amsterdam", graph.Nodes[2].name)
	assert.Equal(t, "Istanbul", graph.Nodes[9].name)

	graph.AddEdge(nodes["London"], nodes["Paris"], 80)
	graph.AddEdge(nodes["London"], nodes["Luxembourg"], 75)
}

func TestAddEdge(t *testing.T) {
	graph := NewGrapth()

	nodes := graph.AddNodes([]string{
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
	})

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

	assert.Equal(t, "Paris", graph.Edges["London"][0].node.name)
	assert.Equal(t, "Luxembourg", graph.Edges["London"][1].node.name)
	assert.Equal(t, "Amsterdam", graph.Edges["London"][2].node.name)

	assert.Equal(t, "London", graph.Edges["Paris"][0].node.name)

	assert.Equal(t, "London", graph.Edges["Luxembourg"][0].node.name)
	assert.Equal(t, "Paris", graph.Edges["Luxembourg"][1].node.name)
	assert.Equal(t, "Berlin", graph.Edges["Luxembourg"][2].node.name)

	assert.Equal(t, "Luxembourg", graph.Edges["Berlin"][0].node.name)
	assert.Equal(t, "Zurich", graph.Edges["Berlin"][1].node.name)
	assert.Equal(t, "Amsterdam", graph.Edges["Berlin"][2].node.name)
	assert.Equal(t, "Vienna", graph.Edges["Berlin"][3].node.name)
}

func TestDijkstra(t *testing.T) {
	graph := NewGrapth()

	nodes := graph.AddNodes([]string{
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
	})

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

	toCity := "Istanbul"
	Dijkstra(graph, "London", toCity)

	endNode := graph.GetNode(toCity)

	for n := endNode; n.through != nil; n = n.through {
		fmt.Print(n.name, " <- ")
	}
	fmt.Println("Lodon")
}
