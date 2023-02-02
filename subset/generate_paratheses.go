package subset

// Time complexity: https://www.youtube.com/watch?v=eoofvKI_Okg

func backtracking(cntOpen, cntClose, n int, curComb []rune, combs *[][]rune) {
	if cntOpen >= n && cntClose >= n {
		tmp := make([]rune, len(curComb))
		copy(tmp, curComb)
		*combs = append(*combs, tmp)
		return
	}

	if cntOpen < n {
		curComb = append(curComb, '(')
		backtracking(cntOpen+1, cntClose, n, curComb, combs)
		curComb = curComb[:len(curComb)-1]
	}

	if cntClose < cntOpen {
		curComb = append(curComb, ')')
		backtracking(cntOpen, cntClose+1, n, curComb, combs)
		curComb = curComb[:len(curComb)-1]
	}

}

// generateCombinations function is used to generate all
// braces combinartions
func generateCombinations(n int) [][]rune {
	result := [][]rune{}
	curComb := []rune{}
	backtracking(0, 0, n, curComb, &result)
	return result
}
