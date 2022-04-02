package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveMakeValid(t *testing.T) {
	type DataTest struct {
		input          string
		expectedOutput string
	}

	tests := []DataTest{
		{"", ""},
		{"abc", "abc"},
		{"))(()", "()"},
		{"(abc", "abc"},
		{"duydep(traivl", "duydeptraivl"},
		{"duydep(tr(a)ivl", "duydeptr(a)ivl"},
		{"duydep(tr(a)i)vl", "duydep(tr(a)i)vl"},
		{"duy(dep(tr(a)i)vl", "duydep(tr(a)i)vl"},
	}
	for _, test := range tests {
		realOutput := minRemoveMakeValid(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}

}

func TestGetInvalid(t *testing.T) {
	type DataTest struct {
		input          string
		expectedOutput []int
	}

	tests := []DataTest{
		{"(abc", []int{0}},
		{"))((", []int{0, 1, 2, 3}},
		{"))(()", []int{0, 1, 2}},
		{"))(())", []int{0, 1}},
		{"duydep(traivl", []int{6}},
		{"duydep(tr(a)ivl", []int{6}},
	}
	for _, test := range tests {
		realOutput := getInvalid(test.input)
		assert.Equal(t, test.expectedOutput, realOutput)
	}
}
