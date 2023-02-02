package permutation

// Time complexity: O(n! * n)
// Space complexity: O(n!)
func permutationRecursive(idx int, nums []int, pers *[][]int) {
	// base case
	if idx == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*pers = append(*pers, tmp)
	}

	for i := idx; i < len(nums); i++ {
		nums[i], nums[idx] = nums[idx], nums[i]
		permutationRecursive(idx+1, nums, pers)
		nums[i], nums[idx] = nums[idx], nums[i]
	}

}

func permutation(nums []int) [][]int {
	result := [][]int{}
	permutationRecursive(0, nums, &result)
	return result
}
