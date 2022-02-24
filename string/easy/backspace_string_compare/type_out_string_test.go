package backspacecompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeOut(t *testing.T) {
	type DataTest struct {
		input          string
		expectedOutput string
	}

	tests := []DataTest{
		{
			input:          "abc#",
			expectedOutput: "ab",
		},
		{
			input:          "abc##",
			expectedOutput: "a",
		},
		{
			input:          "abc####",
			expectedOutput: "",
		},
		{
			input:          "abc#####",
			expectedOutput: "",
		},
		{
			input:          "abc#####d",
			expectedOutput: "d",
		},
		{
			input:          "abc#####dee#",
			expectedOutput: "de",
		},
	}

	for _, test := range tests {
		realOutput := typeOut(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}

func TestCheckEqualTypeOutString(t *testing.T) {
	type DataTest struct {
		inputStr1      string
		inputStr2      string
		expectedOutput bool
	}

	tests := []DataTest{
		{
			inputStr1:      "ab#",
			inputStr2:      "a",
			expectedOutput: true,
		},
		{
			inputStr1:      "a#",
			inputStr2:      "",
			expectedOutput: true,
		},
		{
			inputStr1:      "a####",
			inputStr2:      "",
			expectedOutput: true,
		},
		{
			inputStr1:      "abc#de#",
			inputStr2:      "abd",
			expectedOutput: true,
		},
		{
			inputStr1:      "a##e",
			inputStr2:      "e",
			expectedOutput: true,
		},
		{
			inputStr1:      "####a##e",
			inputStr2:      "e",
			expectedOutput: true,
		},
		{
			inputStr1:      "bxj##tw",
			inputStr2:      "bxo#j##tw",
			expectedOutput: true,
		},
		{
			inputStr1:      "bxj##tw",
			inputStr2:      "bxj###tw",
			expectedOutput: false,
		},
		{
			inputStr1:      "a#c",
			inputStr2:      "b",
			expectedOutput: false,
		},
	}

	for _, test := range tests {
		//realOutput := backspaceCompare(test.inputStr1, test.inputStr2)
		//assert.Equal(t, test.expectedOutput, realOutput)

		realOutput1 := backspaceCompare2Pointer(test.inputStr1, test.inputStr2)
		assert.Equal(t, test.expectedOutput, realOutput1)
	}
}
