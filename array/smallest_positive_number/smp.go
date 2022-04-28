package smallestpositivenumber

import (
	"fmt"
	"strconv"
)

func Solution(a []int) int {
	smallestMissingPosNum := findSmallestMissingPositiveInteger(a)
	return getNumberOf1(smallestMissingPosNum)
}

func getNumberOf1(num int) int {
	numInBinary := strconv.FormatInt(int64(num), 2)
	count := 0
	for i := 0; i < len(numInBinary); i++ {
		if numInBinary[i] == '1' {
			count++
		}
	}
	return count
}

func findSmallestMissingPositiveInteger(a []int) int {
	mapNumToIn := make(map[int]int)

	hashPositive := false
	for i, num := range a {
		if num >= 1 {
			hashPositive = true
		}
		mapNumToIn[num] = i
	}

	if !hashPositive {
		return 1
	}

	for i := 1; i <= len(a)+1; i++ {
		fmt.Println(i)
		if _, ok := mapNumToIn[i]; !ok {
			return i
		}
	}
	return len(a) + 1

}
