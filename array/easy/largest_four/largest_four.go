package largestfour

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sort(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}

func largestFour(arr []int32) int32 {
	// Write your code here
	//

	if len(arr) <= 4 {
		sum := int32(0)
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		return sum
	}

	temp := sort(arr[:4])

	max1 := temp[0]
	max2 := temp[1]
	max3 := temp[2]
	max4 := temp[3]

	for i := 4; i < len(arr); i++ {
		if arr[i] > max1 {
			max4 = max3
			max3 = max2
			max2 = max1
			max1 = arr[i]
		} else if arr[i] > max2 {
			max4 = max3
			max3 = max2
			max2 = arr[i]
		} else if arr[i] > max3 {
			max4 = max3
			max3 = arr[i]
		} else if arr[i] > max4 {
			max4 = arr[i]
		}
	}

	return max1 + max2 + max3 + max4
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := largestFour(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
