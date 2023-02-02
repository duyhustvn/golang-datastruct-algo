package subset

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// the order is not important
func compare2DSlices(sl1 [][]rune, sl2 [][]rune) bool {
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

func TestGenerateParatheses(t *testing.T) {
	type TestData struct {
		input          int
		expectedOutput [][]rune
	}

	tests := []TestData{
		{
			1, [][]rune{{'(', ')'}},
		},
		{
			2, [][]rune{{'(', ')', '(', ')'}, {'(', '(', ')', ')'}},
		},
	}

	for _, test := range tests {
		actualOutput := generateCombinations(test.input)
		assert.Equal(t, true, compare2DSlices(actualOutput, test.expectedOutput))
	}
}
