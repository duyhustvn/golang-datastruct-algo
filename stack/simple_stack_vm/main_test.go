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
		{[]string{}, []int{}},
		{[]string{"PUSH a"}, []int{}},

		{[]string{"PUSH 10"}, []int{10}},
		{[]string{"PUSH 10", "PUSH 20"}, []int{20, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 10"}, []int{10, 20, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 10", "PUSH 40"}, []int{40, 10, 20, 10}},

		{[]string{"MUL"}, []int{}},
		{[]string{"MUL 10"}, []int{}},
		{[]string{"PUSH 10", "MUL"}, []int{}},
		{[]string{"PUSH 10", "MUL 10"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "MUL"}, []int{200}},
		{[]string{"PUSH 10", "PUSH 20", "MUL 10"}, []int{200}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "MUL"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 10"}, []int{10, 200}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 10", "MUL"}, []int{2000}},

		{[]string{"ADD"}, []int{}},
		{[]string{"PUSH 10", "ADD"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "ADD"}, []int{30}},
		{[]string{"PUSH 10", "PUSH 20", "ADD", "ADD"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "ADD", "PUSH 10"}, []int{10, 30}},
		{[]string{"PUSH 10", "PUSH 20", "ADD", "PUSH 10", "ADD"}, []int{40}},

		{[]string{"SWAP"}, []int{}},
		{[]string{"PUSH 10", "SWAP"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "SWAP"}, []int{10, 20}},
		{[]string{"PUSH 10", "PUSH 20", "SWAP", "PUSH 30"}, []int{30, 10, 20}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 30", "SWAP"}, []int{20, 30, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 30", "SWAP", "SWAP"}, []int{30, 20, 10}},
		{[]string{"PUSH 10", "PUSH 20", "PUSH 30", "SWAP", "SWAP", "SWAP"}, []int{20, 30, 10}},

		{[]string{"PUSH 20", "SWAP", "ADD"}, []int{}},
		{[]string{"PUSH 10", "PUSH 20", "SWAP", "ADD"}, []int{30}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 7", "ADD"}, []int{207}},
		{[]string{"PUSH 10", "PUSH 20", "MUL", "PUSH 7", "ADD", "ADD"}, []int{}},
		{[]string{"PUSH 10", "PUSH 0", "MUL", "PUSH 7", "ADD"}, []int{7}},
	}
	for _, test := range tests {
		realOutput := Solution(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}
