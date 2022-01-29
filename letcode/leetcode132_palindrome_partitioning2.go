package letcode

// 思路: 得到每个 i 对应的值，i + 1 下的值，只需要遍历所有的 i，得到一个值即可
func minCut(s string) int {

	// 得到所有的回文子串
	palin := make([][]bool, len(s), len(s))
	// 当 j < i 时设置为 true
	for i := 0; i < len(s); i++ {
		palin[i] = make([]bool, len(s))
		for j := 0; j <= i; j++ {
			palin[i][j] = true
		}
	}
	// j > j 时候
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			palin[i][j] = palin[i + 1][j - 1] && s[i] == s[j]
		}
	}

	if len(s) <= 1 {
		return 0
	}
	result := make([]int, len(s))
	result[0] = 0

	// a b b a
	for i := 1; i < len(s); i++ {
		// 如果整个数组是回文的
		if palin[0][i] {
			result[i] = 0
		} else {
			// 到第 i 个点最多切 i 刀
			result[i] = i
			for j := 0; j < i; j++ {
				// 如果是回文的, 而且变小了
				if result[j] + 1 < result[i] && palin[j + 1][i] {
					result[i] = result[j] + 1
				}
			}
		}
	}
	return result[len(s) - 1]

}
