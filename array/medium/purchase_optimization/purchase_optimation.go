package purchaseoptimization

func sumToPos(prices []int32) []int64 {
	priceSum := make([]int64, len(prices)+1)
	for i := 0; i < len(prices); i++ {
		priceSum[i+1] = int64(prices[i]) + priceSum[i]
	}

	return priceSum
}

func upperBound(array []int64, target int64) int32 {
	low, high, mid := int32(0), int32(len(array)-1), int32(0)

	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func findMaximumValue(prices []int32, pos []int32, amount []int64) []int32 {
	// Write your code here

	count := make([]int32, len(pos))

	// priceSumTilPos[i] is the sum of first i item in prices
	priceSumTilPos := sumToPos(prices)

	for i := 0; i < len(pos); i++ {
		startPos := pos[i] - 1

		finalPos := upperBound(priceSumTilPos, priceSumTilPos[startPos]+amount[i]) - 1
		count[i] = finalPos - startPos
	}
	return count

}
