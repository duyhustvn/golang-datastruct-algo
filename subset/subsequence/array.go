package subsequence

func sequencesWithFixedLength(arr []int, length int, idx int, subsequence []int, subsequences *[][]int) {
	if len(subsequence) > length {
		return
	}

	if idx == len(arr) {
		if len(subsequence) == length {
			*subsequences = append(*subsequences, subsequence)
		}
		return
	}

	subsequence = append(subsequence, arr[idx])
	sequencesWithFixedLength(arr, length, idx+1, subsequence, subsequences)

	// not pick element to subsequence
	subsequence = subsequence[:len(subsequence)-1]
	sequencesWithFixedLength(arr, length, idx+1, subsequence, subsequences)
}

// Generate all sub sequences
// order is not important
// Time complexity: O(n2^n) n is number of copy time
// Space complexity: O(n) we use O(n) to maintain subsequence and modify subsequence in-place. 
func allSequenceRecursive(arr []int, idx int, subsequence []int, subsequences *[][]int) {
	if idx == len(arr) {
		*subsequences = append(*subsequences, subsequence)
		return
	}

	// pick to subsequence
	subsequence = append(subsequence, arr[idx])
	allSequenceRecursive(arr, idx+1, subsequence, subsequences)

	// not pick to subsequence
	subsequence = subsequence[:len(subsequence)-1]
	allSequenceRecursive(arr, idx+1, subsequence, subsequences)
}

// Time complexity: O(n2^n) n is number of copy time
// Space complexity: O(n2^n) n is number of element in each subset
func allSequenceIterative(arr []int) [][]int {
	n := len(arr)

	// sequences := make([][]int, math.Pow(2, n))
	sequences := [][]int{}
	sequences = append(sequences, []int{})
	for i := 0; i < n; i++ {
		num := arr[i]

		tmp := make([][]int, len(sequences))
		copy(tmp, sequences)
		for j := 0; j < len(tmp); j++ {
			tmp[j] = append(tmp[j], num)
		}

		sequences = append(sequences, tmp...)
	}
	return sequences
}

func generateSubsequenceIterative(arr []int) [][]int {
	return allSequenceIterative(arr)
}

// generate all subset
func generateSubsequenceRecursive(arr []int) [][]int {
	subsequence := []int{}
	subsequences := [][]int{}

	allSequenceRecursive(arr, 0, subsequence, &subsequences)
	return subsequences
}

func generateSubsequenceWithFixedLength(arr []int, length int) [][]int {
	subsequence := []int{}
	subsequences := [][]int{}

	sequencesWithFixedLength(arr, length, 0, subsequence, &subsequences)
	return subsequences
}
