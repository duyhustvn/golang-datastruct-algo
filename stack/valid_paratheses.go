package stack

/*
  solution for problem at leetcode https://leetcode.com/problems/valid-parentheses/
  @param s input string for checking
*/
func isValid(s string) bool {
	stack := Stack{}
	symbolMap := map[uint8]uint8{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack.Push(s[i])
		} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			symbol := stack.Pop()
			if symbol == nil {
				return false
			}
			if symbol.(uint8) != symbolMap[s[i]] {
				return false
			}
		} else {
			return false
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}
