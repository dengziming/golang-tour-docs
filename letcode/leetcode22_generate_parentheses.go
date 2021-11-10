package letcode

// 动态规则，递归其实更简单，但是会重复计算
func generateParenthesis(n int) []string {
	len := n + 1
	result := make([][][]string, len)

	for i := 0; i < len; i++ {
		result[i] = make([][]string, len)
	}

	result[0][0] = []string{""}

	for i := 1; i < len; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				result[i][j] = appendSuffix(result[i - 1][j], "(")
			} else if i == j {
				result[i][j] = appendSuffix(result[i][j - 1], ")")
			} else {
				result[i][j] = append(appendSuffix(result[i - 1][j], "("), appendSuffix(result[i][j - 1], ")")...)
			}
		}
	}
	return result[len- 1][len- 1]
}

func appendSuffix(strs []string, suffix string) [] string {
	result := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		result[i] = strs[i] + suffix
	}
	return result
}
