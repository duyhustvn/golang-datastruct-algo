package stack

import (
	"sort"
)

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

type PositionToValue struct {
	pos   int
	value uint8
}

func minRemoveMakeValid(s string) string {
	invalids := getInvalid(s)
	for i, invalid := range invalids {
		s = s[:invalid-i] + s[invalid-i+1:]
	}
	return s
}

func getInvalid(s string) []int {
	stack := Stack{}
	invalidParathesesIndex := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			pos2Value := PositionToValue{i, s[i]}
			stack.Push(pos2Value)
			continue
		} else if s[i] == ')' {
			if stack.IsEmpty() {
				invalidParathesesIndex = append(invalidParathesesIndex, i)
				continue
			} else {
				stack.Pop()
			}
		}
	}

	for !stack.IsEmpty() {
		invalid := stack.Pop()
		invalidParathesesIndex = append(invalidParathesesIndex, invalid.(PositionToValue).pos)
	}
	sort.Slice(invalidParathesesIndex, func(i, j int) bool {
		return invalidParathesesIndex[i] < invalidParathesesIndex[j]
	})

	return invalidParathesesIndex
}
