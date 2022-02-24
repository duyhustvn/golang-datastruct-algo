package trappingrainwater

/*
https://leetcode.com/problems/trapping-rain-water/
  **/

/*
   Find max value in the slice
  **/
func maxInSlice(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func TrappingWaterBruteForce(nums []int) int {
	totalWater := 0
	for i := range nums {
		maxL := maxInSlice(nums[:i])
		maxR := maxInSlice(nums[i:])
		min := maxL
		if maxL > maxR {
			min = maxR
		}
		currentWater := min - nums[i]
		if currentWater >= 0 {
			totalWater += currentWater
		}
	}
	return totalWater
}

/*
  Approach: Two pointer
   Initialize left pointer to 0 and right pointer to size -1
   while left < right do:
     if nums[left] is smaller than nums[right]
       if nums[left] >= max_left, update max_left
       else add left_max - nums[i] to ans
       add 1 to left
     else
       if nums[right] is smaller than nums[left]
         if nums[right] >= max_right, update max_right
         else add max_right - nums[j] to ans
         subtract 1 from right
  **/

func TrappingWaterOptimization(nums []int) int {
	i, j := 0, len(nums)-1
	totalWater := 0
	maxL, maxR := 0, 0

	for i < j {
		if nums[i] <= nums[j] {
			if nums[i] >= maxL {
				maxL = nums[i]
			} else {
				totalWater += maxL - nums[i]
			}
			i++
		} else {
			if nums[j] >= maxR {
				maxR = nums[j]
			} else {
				totalWater += maxR - nums[j]
			}
			j--
		}
	}
	return totalWater
}
