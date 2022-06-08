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
		{[][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, map[int][]int{0: {1}, 1: {2}, 5: {2, 3, 4}, 3: {0, 4}}},
	}

	for _, test := range tests {
		actual := buildAdjacencyList(test.prerequisites)
		// t.Log(actual)
		assert.Equal(t, test.expected[0], actual[0])
		assert.Equal(t, test.expected[1], actual[1])
		assert.Equal(t, test.expected[5], actual[5])
		assert.Equal(t, test.expected[3], actual[3])
	}
}
