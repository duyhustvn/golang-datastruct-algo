package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	type DataTest struct {
		input          []int
		expectedOutput []int
	}

	tests := []DataTest{
		{[]int{1}, []int{1}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{1, 0, 5, 4}, []int{0, 1, 4, 5}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		BubbleSort(test.input)
		assert.Equal(t, test.expectedOutput, test.input)
	}
}
