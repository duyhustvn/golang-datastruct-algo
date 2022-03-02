package lengthoflongestsubstring

func lengthOfLongestSubstring(str string) int {
	mapChar2Id := make(map[string]int)
	max := 0

	for i := 0; i < len(str); i++ {
		foundIndex, ok := mapChar2Id[string(str[i])]
		if ok {
			if len(mapChar2Id) > max {
				max = len(mapChar2Id)
			}
			i = foundIndex
			mapChar2Id = make(map[string]int)
		} else {
			mapChar2Id[string(str[i])] = i
			if i == len(str)-1 {
				if len(mapChar2Id) > max {
					max = len(mapChar2Id)
				}
			}
		}
	}

	return max
}

func LenghOfLongestSubstringSlideWindow(str string) int {
	left, right := 0, 0
	seen := make(map[string]int)
	max := 0

	for right <= len(str)-1 {
		foundIndex, ok := seen[string(str[right])]
		if !ok {
			seen[string(str[right])] = right
			right++
			if right-left > max {
				max = right - left
			}
		} else {
			if foundIndex < left {
				seen[string(str[foundIndex])] = right
				right++
				if right-left > max {
					max = right - left
				}
			} else {
				left = foundIndex + 1
			}
		}
	}
	return max
}
