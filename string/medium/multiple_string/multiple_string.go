package multiplestring

import (
	"math"
	"strconv"
)

func ConvertStringToNumber(str string) int {
	num := 0
	totalNumber := len(str) - 1
	for i := len(str) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(str[i]))
		num += int(math.Pow(10, float64(totalNumber-i))) * n
	}
	return num
}

func ConvertNumberToString(num int64, str string) string {
	if num != 0 {
		i := num % 10
		str = string(i) + str
		return ConvertNumberToString(num/10, str)
	}
	return str
}
