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
