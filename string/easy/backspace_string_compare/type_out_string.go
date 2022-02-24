package backspacecompare

func backspaceCompare(str1 string, str2 string) bool {
	typeOutStr1 := typeOut(str1)
	typeOutStr2 := typeOut(str2)
	return typeOutStr1 == typeOutStr2
}

/*
   Remove character before #

   abc#d -> abd
   abc##d -> ad
  **/
func typeOut(strs string) string {
	var stack string
	for _, str := range strs {
		if string(str) == "#" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				stack = ""
			}
		} else {
			stack += string(str)
		}
	}
	return stack
}

// Optimization solution
func backspaceCompare2Pointer(str1 string, str2 string) bool {
	i, j, numberHashStr1, numberHashStr2 := len(str1)-1, len(str2)-1, 0, 0
	for {
		if i < 0 && j < 0 {
			break
		}
		if i >= 0 && string(str1[i]) == "#" {
			numberHashStr1++
			i--
			continue
		}
		if j >= 0 && string(str2[j]) == "#" {
			numberHashStr2++
			j--
			continue
		}

		if numberHashStr1 > 0 {
			numberHashStr1--
			i--
			continue
		}

		if numberHashStr2 > 0 {
			numberHashStr2--
			j--
			continue
		}

		if i >= 0 && j >= 0 {
			if string(str1[i]) != string(str2[j]) {
				return false
			}
		} else if (i >= 0 && j < 0) || (j >= 0 && i < 0) {
			return false
		}
		i--
		j--
	}

	return true
}
