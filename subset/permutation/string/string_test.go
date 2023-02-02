package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compare2DSlices(sl1 []string, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for _, str1 := range sl1 {
		isFound := false
		for j, str2 := range sl2 {
			if str1 == str2 {
				// delete str2
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
		input                string
		expectedPermutations []string
	}

	tests := []DataTest{
		{
			"a", []string{"a"},
		},
		{
			"ab", []string{"ab", "ba"},
		},
		{
			"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}

	for _, test := range tests {
		realPermutations := permutation(test.input)
		equal := compare2DSlices(realPermutations, test.expectedPermutations)

		if !equal {
			for _, str := range realPermutations {
				fmt.Println(str)
			}
		}

		assert.Equal(t, equal, true)
	}
}
