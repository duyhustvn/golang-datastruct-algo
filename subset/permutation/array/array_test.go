package permutation

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// the order is not important
func compare2DSlices(sl1 [][]int, sl2 [][]int) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for _, sub1 := range sl1 {
		isFound := false
		for j, sub2 := range sl2 {
			if reflect.DeepEqual(sub1, sub2) {
				// remove sub2 from sl2
				copy(sl2[j:], sl2[j+1:])
				sl2 = sl2[:len(sl2)-1]
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}

	return true
}

func TestPermutation(t *testing.T) {
	type DataTest struct {
		input                []int
		expectedPermutations [][]int
	}

	tests := []DataTest{
		{
			[]int{1, 2}, [][]int{{1, 2}, {2, 1}},
		},
		{
			[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 3, 2}, {1, 4, 2, 3}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 3, 2, 1}, {4, 3, 1, 2}, {4, 1, 3, 2}, {4, 1, 2, 3}},
		},
	}

	for _, test := range tests {
		realPermutations := permutation(test.input)

		assert.Equal(t, compare2DSlices(realPermutations, test.expectedPermutations), true)
	}
}
