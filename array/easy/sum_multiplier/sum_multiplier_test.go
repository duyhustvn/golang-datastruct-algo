package summultiplier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumMultiplier(t *testing.T) {
	type DataTest struct {
		input    []int32
		expected bool
	}

	tests := []DataTest{
		{[]int32{-1, -5, 0, 11, -5}, true},
	}

	for _, test := range tests {
		actual := sumMultiplier(test.input)
		assert.Equal(t, test.expected, actual)
	}

}
