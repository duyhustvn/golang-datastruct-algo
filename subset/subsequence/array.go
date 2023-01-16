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

// generate all subset
func generateSubsequence(arr []int) [][]int {
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
