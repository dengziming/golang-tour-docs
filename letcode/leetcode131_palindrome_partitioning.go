package letcode

// 思路1: 递归。每次任务取一段，只要是回文就继续，直到结束
// 思路2: 回溯，由于递归写太多遍了，这次换回溯
func partition(s string) [][]string {
	result := make([][]string, 0)

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

	dfs(&palin, s, 0, make([]string, 0), &result)

	return result
}

func dfs(palin *[][]bool, s string, index int, ans []string, result *[][]string) {
	// 递归机
	if index == len(s) {
		*result = append(*result, append([]string(nil), ans...))
		return
	}
	for i := index; i < len(s); i++ {
		if (*palin)[index][i] {
			dfs(palin, s, i + 1, append(ans, s[index: i + 1]), result)
			print()
		}
	}
}
