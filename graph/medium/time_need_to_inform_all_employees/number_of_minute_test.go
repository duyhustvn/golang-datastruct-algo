package time_need_to_inform_all_employees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAdjacencyList(t *testing.T) {
	type TestData struct {
		headerId int
		managers []int
	}

	tests := []TestData{
		{6, []int{2, 2, 6, 4, 6, 4, -1, 6, 6}},
	}

	for _, test := range tests {
		actualOutput := buildAdjacencyList(test.headerId, test.managers)
		t.Log(actualOutput)
		suborsOf6, ok := actualOutput[6]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{2, 4, 7, 8}, suborsOf6)

		suborsOf2, ok := actualOutput[2]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{0, 1}, suborsOf2)

		suborsOf4, ok := actualOutput[4]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{3, 5}, suborsOf4)
	}
}

func TestBuildGraphFromAdjacencyList(t *testing.T) {
	type TestData struct {
		adjacencyList map[int][]int
		expectedDepth int
	}

	tests := []TestData{
		{map[int][]int{6: []int{2, 4, 7, 8}}, 2},
		{map[int][]int{6: []int{2, 4, 7, 8}, 2: []int{0, 1}}, 3},
	}

	for _, test := range tests {
		graph, headNode := buildGraphFromAdjacencyList(test.adjacencyList, 6)
		graph.String()
		assert.NotNil(t, headNode)

	}
}

func TestGetDepth(t *testing.T) {
	type TestData struct {
		adjacencyList map[int][]int
		informTime    []int
		expectedDepth int
	}

	tests := []TestData{
		{map[int][]int{6: []int{2, 4, 7, 8}, 2: []int{0, 1}}, []int{0, 0, 2, 0, 1, 0, 2, 0, 0}, 4},
	}

	for _, test := range tests {
		graph, headNode := buildGraphFromAdjacencyList(test.adjacencyList, 6)
		graph.String()
		assert.NotNil(t, headNode)

		actualDepth := getDepth(graph, headNode, test.informTime)
		assert.Equal(t, test.expectedDepth, actualDepth)
	}
}

func TestNumberOfMinutes(t *testing.T) {
	type TestData struct {
		n            int
		headID       int
		manager      []int
		informTime   []int
		expectedTime int
	}

	tests := []TestData{
		{9, 6, []int{2, 2, 6, 4, 6, 4, -1, 6, 6}, []int{0, 0, 2, 0, 1, 0, 2, 0, 0}, 4},
		{1, 0, []int{-1}, []int{0}, 0},
	}

	for _, test := range tests {
		actualTime := numOfMinutes(test.n, test.headID, test.manager, test.informTime)
		assert.Equal(t, test.expectedTime, actualTime)

		actualTime1 := numOfMinutesWeightedGraph(test.n, test.headID, test.manager, test.informTime)
		assert.Equal(t, test.expectedTime, actualTime1)

		actualTime2 := numOfMinutesOptimal(test.n, test.headID, test.manager, test.informTime)
		assert.Equal(t, test.expectedTime, actualTime2)
	}
}

func TestBuildAdjacencyListWithWeight(t *testing.T) {
	type TestData struct {
		headerId   int
		managers   []int
		informTime []int
	}

	tests := []TestData{
		{6, []int{1, 6, 1, 4, 6, 4, -1, 6, 6}, []int{0, 1, 0, 0, 2, 0, 3, 0, 0}},
	}

	for _, test := range tests {
		actualOutput := buildAdjacencyListWithWeight(test.managers, test.informTime)
		t.Log(actualOutput)
		suborsOf6, ok := actualOutput[6]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{1, 4, 7, 8}, suborsOf6.Sub)
		assert.Equal(t, 3, suborsOf6.Weight)

		suborsOf4, ok := actualOutput[4]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{3, 5}, suborsOf4.Sub)
		assert.Equal(t, 2, suborsOf4.Weight)

		suborsOf1, ok := actualOutput[1]
		assert.Equal(t, true, ok)
		assert.Equal(t, []int{0, 2}, suborsOf1.Sub)
		assert.Equal(t, 1, suborsOf1.Weight)
	}
}

func TestBuildWeightGraphFromAdjacencyList(t *testing.T) {
	type TestData struct {
		headerId      int
		adjacencyList map[int]AdjacencyListWithWeight
	}

	tests := []TestData{
		{6, map[int]AdjacencyListWithWeight{
			1: {[]int{0, 2}, 1},
			4: {[]int{3, 5}, 2},
			6: {[]int{1, 4, 7, 8}, 3},
		}},
	}

	for _, test := range tests {
		graph, headNode := buildWeightGraphFromAdjacencyList(test.headerId, test.adjacencyList)
		graph.String()
		assert.NotNil(t, graph)
		assert.NotNil(t, headNode)

	}

}
