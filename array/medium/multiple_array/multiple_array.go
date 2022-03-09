package multiplearray

func MulttipleArray(arr1 []int, arr2 []int) []int {
	mapMultiple := make(map[int][]int)

	total := []int{}

	for i := len(arr2) - 1; i >= 0; i-- {
		res, ok := mapMultiple[arr2[i]]
		if !ok {
			res = MultipleNumber(arr1, arr2[i], 0, []int{})
			mapMultiple[arr2[i]] = res
		}

		for j := 0; j < len(arr2)-1-i; j++ {
			res = append(res, 0)
		}

		if len(total) > len(res) {
			total = Sum2Array(total, res)
		} else {
			total = Sum2Array(res, total)
		}
	}
	return total
}

func MultipleNumber(arr1 []int, mul int, carrier int, result []int) []int {
	if len(arr1) > 0 {
		r := arr1[len(arr1)-1]*mul + carrier
		carrier = r / 10
		i := r % 10
		result = append([]int{i}, result...)
		arr1 = arr1[:len(arr1)-1]
		return MultipleNumber(arr1, mul, carrier, result)
	} else if len(arr1) == 0 {
		if carrier > 0 {
			result = append([]int{carrier}, result...)
		}
	}
	return result
}

// arr1 len always greater than arr2
func Sum2Array(arr1 []int, arr2 []int) []int {
	result := []int{}
	carrier := 0
	j := len(arr2) - 1
	for i := len(arr1) - 1; i >= 0; i-- {
		two := 0
		if j >= 0 {
			two = arr2[j]
		}
		j--
		rs := arr1[i] + two + carrier
		carrier = rs / 10
		addedNum := rs % 10
		result = append([]int{addedNum}, result...)
		if i == 0 && carrier > 0 {
			result = append([]int{carrier}, result...)
		}
	}
	return result
}
