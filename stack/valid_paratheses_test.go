package stack

import "testing"

type Tests struct {
	Input          string `json:"input"`
	ExpectedOutput bool   `json:"expected_output"`
}

func TestValidParatheses(t *testing.T) {
	tests := []Tests{
		{
			Input:          "",
			ExpectedOutput: true,
		},
		{
			Input:          "()",
			ExpectedOutput: true,
		},
		{
			Input:          "([])",
			ExpectedOutput: true,
		},
		{
			Input:          "([]{})",
			ExpectedOutput: true,
		},
		{
			Input:          "(){}[]",
			ExpectedOutput: true,
		},
		{
			Input:          "([]{}[{}])[]{}()[]",
			ExpectedOutput: true,
		},
		{
			Input:          "(]",
			ExpectedOutput: false,
		},
		{
			Input:          "(",
			ExpectedOutput: false,
		},
		{
			Input:          "([]",
			ExpectedOutput: false,
		},
		{
			Input:          "([",
			ExpectedOutput: false,
		},
		{
			Input:          "()]",
			ExpectedOutput: false,
		},
	}

	for _, test := range tests {
		realOutput := isValid(test.Input)
		if realOutput != test.ExpectedOutput {
			t.Errorf("Expect %+v insted of %+v for %+v", test.ExpectedOutput, realOutput, test.Input)
		}
	}
}
