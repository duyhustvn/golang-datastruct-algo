package validpalindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type DataTest struct {
		input          string
		expectedOutput bool
	}

	tests := []DataTest{
		{"", true},
		{"aba", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"0P", false},
	}

	for _, test := range tests {
		realOutput := isPalindrome(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}

func TestIsAlmostPalindrome(t *testing.T) {
	type DataTest struct {
		input          string
		expectedOutput bool
	}

	tests := []DataTest{
		{"", true},
		{"aba", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", true},
		{"0P", true},
		{"ab", true},
	}

	for _, test := range tests {
		realOutput := isAlmostPalindrome(test.input, false)
		if !assert.Equal(t, test.expectedOutput, realOutput) {
			fmt.Println(test.input)
		}
	}
}
