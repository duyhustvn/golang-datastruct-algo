package validpalindrome

// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	//fmt.Println(s)
	left, right := 0, len(s)-1
	var charLeft, charRight uint8
	for left <= right {
		if ('a' <= s[left] && s[left] <= 'z') || ('0' <= s[left] && s[left] <= '9') {
			charLeft = s[left]
		} else if 'A' <= s[left] && s[left] <= 'Z' {
			charLeft = s[left] - 'A' + 'a'
		} else {
			left++
			continue
		}

		if ('a' <= s[right] && s[right] <= 'z') || ('0' <= s[right] && s[right] <= '9') {
			charRight = s[right]
		} else if 'A' <= s[right] && s[right] <= 'Z' {
			charRight = s[right] - 'A' + 'a'
		} else {
			right--
			continue
		}

		//fmt.Printf("left: %d right: %d charLeft: %d charRight: %d\n", left, right, charLeft, charRight)
		if charLeft != charRight {
			return false
		}
		left++
		right--
	}

	return true
}
