package bs

func upperBound(nums []int, low, high, target int) int {
	tmp := high
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low <= tmp && nums[low] <= target {
		low++
	}

	return low
}