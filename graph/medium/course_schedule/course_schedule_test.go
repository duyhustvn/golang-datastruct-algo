package course_schedule

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAdjacencyList(t *testing.T) {
	type TestData struct {
		prerequisites    [][]int
		expectedAdjList  map[int][]int
		expectedIndegree map[int]int
	}

	tests := []TestData{
		{[][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, map[int][]int{0: {1}, 1: {2}, 3: {0, 4}, 5: {2, 3, 4}}, map[int]int{0: 1, 1: 1, 2: 2, 3: 1, 4: 2}},
		{[][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {5, 6}, {6, 4}}, map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, map[int]int{0: 1, 1: 1, 2: 1, 4: 1, 5: 1, 6: 1}},
	}

	for _, test := range tests {
		actualAdjList, actualIndegree := buildAdjacencyList(test.prerequisites)
		// t.Log(actualIndegree)

		assert.Equal(t, true, reflect.DeepEqual(test.expectedAdjList, actualAdjList))
		assert.Equal(t, true, reflect.DeepEqual(test.expectedIndegree, actualIndegree))
	}
}

func TestQueuePush(t *testing.T) {
	q := newQueue()

	var v int
	var ok bool
	v, ok = q.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)
	q.Push(1)
	v, ok = q.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)

	q.Push(1)
	q.Push(2)
	v, ok = q.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)

	v, ok = q.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, v)

	v, ok = q.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)
}

func TestBfs(t *testing.T) {
	type TestData struct {
		adjacency map[int][]int
		fromNode  int
		expected  bool
	}

	tests := []TestData{
		{map[int][]int{0: {1}, 1: {2}, 5: {2, 3, 4}, 3: {0, 4}}, 0, false},
		{map[int][]int{0: {1}, 1: {2}, 5: {2, 3, 4}, 3: {0, 4}}, 1, false},
		{map[int][]int{0: {1}, 1: {2}, 5: {2, 3, 4}, 3: {0, 4}}, 3, false},
		{map[int][]int{0: {1}, 1: {2}, 5: {2, 3, 4}, 3: {0, 4}}, 5, false},

		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 0, false},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 1, false},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 2, false},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 3, false},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 4, true},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 5, true},
		{map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}, 6, true},
	}

	for _, test := range tests {
		actual := bfs(test.fromNode, test.adjacency)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCanFinish(t *testing.T) {
	type TestData struct {
		numberCourses int
		prerequisites [][]int
		expected      bool
	}

	tests := []TestData{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{6, [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, true},
		{7, [][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {5, 6}, {6, 4}}, false},
	}

	for _, test := range tests {
		actual := canFinish(test.numberCourses, test.prerequisites)
		assert.Equal(t, test.expected, actual)
	}
}
