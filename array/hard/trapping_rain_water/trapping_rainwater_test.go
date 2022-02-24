package trappingrainwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Input        []int `json:"input"`
	ExpectedArea int   `json:"expected_area"`
}

func TestTrappingWaterBruteForce(t *testing.T) {
	type TestData struct {
		Input          []int `json:"input"`
		ExpectedOutput int   `json:"expected_output"`
	}

	tests := []TestData{
		{
			Input:          []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
			ExpectedOutput: 8,
		},
		{
			Input:          []int{3, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
			ExpectedOutput: 15,
		},
		{
			Input:          []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			ExpectedOutput: 6,
		},
	}

	for _, test := range tests {
		realOutput := TrappingWaterBruteForce(test.Input)
		assert.Equal(t, test.ExpectedOutput, realOutput)

		realOutput1 := TrappingWaterOptimization(test.Input)
		assert.Equal(t, test.ExpectedOutput, realOutput1)
	}

}
