package main

import (
	"ds/graph/medium/time_need_to_inform_all_employees"
	lengthoflongestsubstring "ds/string/medium/length_of_longest_substring"
	"fmt"
)

func main() {
	lengthoflongestsubstring.LenghOfLongestSubstringSlideWindow("asjrgapa")
	fmt.Println(time_need_to_inform_all_employees.NumOfMinutesOptimal(9, 6, []int{2, 2, 6, 4, 6, 4, -1, 6, 6}, []int{0, 0, 3, 0, 1, 0, 2, 0, 0}))
}
