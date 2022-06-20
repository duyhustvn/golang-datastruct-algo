package purchaseoptimization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumToPos(t *testing.T) {
	type DataTest struct {
		price    []int32
		expected []int32
	}

	tests := []DataTest{
		{[]int32{1, 2, 2, 3, 4}, []int32{0, 1, 3, 5, 8, 12}},
	}

	for _, test := range tests {
		actual := sumToPos(test.price)
		assert.Equal(t, test.expected, actual)
	}

}
