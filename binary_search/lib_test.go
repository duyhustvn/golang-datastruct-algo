package bs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUpperBound(t *testing.T) {
	type target struct {
		target int
		expectedIndex int 
	}

	tests := []struct{
		nums []int
		cases []target
	} {
		{
			nums: []int{1, 10, 10, 10, 20, 20, 30, 40, 50, 60}, 
			cases: []target {
				{target: 0, expectedIndex: 0},
				{target: 1, expectedIndex: 1},
				{target: 5, expectedIndex: 1},
				{target: 10, expectedIndex: 4},
				{target: 11, expectedIndex: 4},
				{target: 20, expectedIndex: 6},
				{target: 29, expectedIndex: 6},
				{target: 30, expectedIndex: 7},
				{target: 40, expectedIndex: 8},
				{target: 50, expectedIndex: 9},
				{target: 60, expectedIndex: 10},
				{target: 61, expectedIndex: 10},
			},
		},
	}

	for _, test := range(tests) {
		for _, target := range(test.cases) {
			actual := upperBound(test.nums, 0, len(test.nums)-1, target.target)
			assert.Equal(t, actual, target.expectedIndex)
		}
	}
}