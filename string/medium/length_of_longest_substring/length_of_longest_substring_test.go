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
		{"abcabcbb", 3},
		{"bb", 1},
		{"abbbdab", 3},
		{"nfpdmpi", 5},
		{"asjrgapa", 6},
	}

	for _, test := range tests {
		realOutput := lengthOfLongestSubstring(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)

		realOutput1 := LenghOfLongestSubstringSlideWindow(test.input)
		assert.Equal(t, test.expectedOutput, realOutput1)
	}
}
