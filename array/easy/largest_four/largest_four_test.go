package largestfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	type DataTest struct {
		input    []int32
		expected []int32
		sum      int32
	}

	tests := []DataTest{
		{[]int32{1, 3, 4, 2, 4, 6, 7}, []int32{7, 6, 4, 4, 3, 2, 1}, 21},
		{[]int32{1, 3, 4}, []int32{4, 3, 1}, 8},
	}

	for _, test := range tests {
		actual := sort(test.input)
		assert.Equal(t, test.expected, actual)

		actualSum := largestFour(test.input)
		assert.Equal(t, test.sum, actualSum)
	}
}
