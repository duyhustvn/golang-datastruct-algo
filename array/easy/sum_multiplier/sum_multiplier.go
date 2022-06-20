package summultiplier

func sumMultiplier(arr []int32) bool {
	// Write your code here
	if len(arr) < 2 {
		return false
	}

	var max1, max2, min1, min2 int32

	if arr[0] > arr[1] {
		max1 = arr[0]
		max2 = arr[1]

		min1 = arr[1]
		min2 = arr[0]
	} else {
		max1 = arr[1]
		max2 = arr[0]

		min1 = arr[0]
		min2 = arr[1]
	}

	sum := int32(0)
	for i := 0; i < len(arr); i++ {
		if i >= 2 {
			if arr[i] > max1 {
				max2 = max1
				max1 = arr[i]
			} else if arr[i] > max2 {
				max2 = arr[i]
			}

			if arr[i] < min1 {
				min2 = min1
				min1 = arr[i]
			} else if arr[i] < min2 {
				min2 = arr[i]
			}
		}

		sum += arr[i]
	}

	return ((max1 * max2) > (2 * sum)) || ((min1 * min2) > (2 * sum))
}

func findMaximumValue(prices []int32, pos []int32, amount []int64) []int32 {
	// Write your code here
	countProd := []int32{}

	for ind, eachPos := range pos {
		count := int32(0)
		sum := int64(0)
		for i := eachPos; i < int32(len(prices)); i++ {
			sum += int64(prices[i])
			if sum < amount[ind] {
				count++
			} else {
				break
			}
		}
		countProd = append(countProd, count)
	}

	return countProd
}
