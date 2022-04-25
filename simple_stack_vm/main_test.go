package simplestackvm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type DataTest struct {
		input          []string
		expectedOutput []int
	}

	tests := []DataTest{
		{[]string{"PUSH 10"}, []int{10}},
		{[]string{"PUSH 10", "PUSH 20"}, []int{20, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 10"}, []int{10, 20, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 10", "PUSH 40"}, []int{40, 10, 20, 10}},

		{[]string{"PUSH 10", "MUL"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "MUL"}, []int{200}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "MUL"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 10"}, []int{10, 200}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 10", "MUL"}, []int{2000}},
	}
	for _, test := range tests {
		realOutput := Solution(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}
