package permutation

func swap(str string, i, j int) string {
	r := []rune(str)

	r[i], r[j] = r[j], r[i]
	return string(r)
}

func permutationRecur(idx int, str string, pers *[]string) {
	if idx == len(str) {
		*pers = append(*pers, str)
		return
	}

	for i := idx; i < len(str); i++ {
		str = swap(str, i, idx)
		permutationRecur(idx+1, str, pers)
		str = swap(str, i, idx)
	}
}

func permutation(str string) []string {
	result := []string{}
	permutationRecur(0, str, &result)
	return result
}
