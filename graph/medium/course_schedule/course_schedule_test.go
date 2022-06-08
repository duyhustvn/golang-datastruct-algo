package course_schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAdjacencyList(t *testing.T) {
	type TestData struct {
		prerequisites [][]int
		expected      map[int][]int
	}

	tests := []TestData{
		{[][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, map[int][]int{0: {1}, 1: {2}, 3: {0, 4}, 5: {2, 3, 4}}},
		{[][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {5, 6}, {6, 4}}, map[int][]int{0: {1}, 1: {2}, 3: {0}, 4: {6}, 5: {4}, 6: {5}}},
	}

	for _, test := range tests {
		actual := buildAdjacencyList(test.prerequisites)
		// t.Log(actual)
		assert.Equal(t, test.expected[0], actual[0])
		assert.Equal(t, test.expected[1], actual[1])
		assert.Equal(t, test.expected[2], actual[2])
		assert.Equal(t, test.expected[3], actual[3])
		assert.Equal(t, test.expected[4], actual[4])
		assert.Equal(t, test.expected[5], actual[5])
		assert.Equal(t, test.expected[6], actual[6])
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
