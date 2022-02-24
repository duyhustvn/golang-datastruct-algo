package twosum

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, sum int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == sum {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumUsingMap(nums []int, sum int) []int {
	idxMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		target := sum - nums[i]
		if idx, ok := idxMap[target]; ok {
			return []int{idx, i}
		}
		idxMap[nums[i]] = i
	}
	return nil
}
