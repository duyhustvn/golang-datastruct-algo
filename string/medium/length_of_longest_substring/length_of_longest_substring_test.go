package lengthoflongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubString(t *testing.T) {
	type TestData struct {
		input          string
		expectedOutput int
	}

	tests := []TestData{
		{"", 0},
		{"abcbda", 4},
		{"abbbda", 3},
		{"abbbdab", 3},
		{"bbbbbb", 1},
		{"abcabcbb", 3},
	}

	for _, test := range tests {
		realOutput := lengthOfLongestSubstring(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}
