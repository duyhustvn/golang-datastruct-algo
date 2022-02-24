package containerwithmostwater

/*
   Problem: https://leetcode.com/problems/container-with-most-water/
*/

func findContainerWithMostWater(nums []int) int {
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			heigh := nums[i]
			if nums[i] > nums[j] {
				heigh = nums[j]
			}
			area := (j - i) * heigh
			if area > max {
				max = area
			}
		}
	}
	return max
}

func findContainerWithMostWaterOptimal(nums []int) int {
	max := 0
	leftIndex := 0
	rightIndex := len(nums) - 1
	for {
		heigh := nums[leftIndex]
		if nums[leftIndex] > nums[rightIndex] {
			heigh = nums[rightIndex]
		}
		area := (rightIndex - leftIndex) * heigh
		if area > max {
			max = area
		}

		if nums[leftIndex] > nums[rightIndex] {
			rightIndex--
		} else {
			leftIndex++
		}

		if leftIndex == rightIndex {
			return max
		}
	}

}
