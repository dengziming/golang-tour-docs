package letcode

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := make([]int, len(s))
	result[0] = 0


	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			result[i] = 0
			continue
		}
		if s[i - 1] == '(' {
			result[i] = maybeOverflow(result, i - 2) + 2
			continue
		}
		// si[i-1: i+1] = '))'
		if i - 1 - result[i-1] == -1 {
			result[i] = 0
			continue
		}
		if s[i - 1 - result[i - 1]] == '(' {
			result[i] = result[i - 1] + 2 + maybeOverflow(result, i - 1 - result[i - 1] - 1)
			continue
		}
		result[i] = 0
	}

	max := 0
	for _, v := range result {
		if max < v {
			max = v
		}
	}
	return max
}

func maybeOverflow(s []int, i int) int {
	if i == -1 {
		return 0
	} else {
		return s[i]
	}
}
