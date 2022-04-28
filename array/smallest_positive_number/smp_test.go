package smallestpositivenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type DataTest struct {
		input          []int
		expectedOutput int
	}

	tests := []DataTest{
		{[]int{1, 3, 2, 6, 4, 1, 2}, 2},
		{[]int{-1, -4}, 1},
		{[]int{}, 1},
	}

	for _, test := range tests {
		actualOutput := Solution(test.input)
		assert.Equal(t, test.expectedOutput, actualOutput)
	}
}
