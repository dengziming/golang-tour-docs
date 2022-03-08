package letcode

import "strings"

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)

	// 每一个竖是否有皇后
	flag1 := make([]int, n)
	// 每一正对线是否有皇后，下标为 i - j + n - 1
	flag2 := make([]int, n * 2 - 1)
	// 每一负是否有皇后，下标为 i + j
	flag3 := make([]int, n * 2 -1)

	var dfs func(target []string, result *[][]string)

	dfs = func(target []string, result *[][]string) {
		if len(target) == n {
			// *result = append(*result, target) 就会指针错乱
			*result = append(*result, append([]string(nil), target...))
			return
		}
		i := len(target)
		for j := 0; j < n; j++ {
			if flag1[j] == 0 && flag2[i - j + n - 1] == 0 && flag3[j + i] == 0 {
				flag1[j] = 1
				flag2[i - j + n - 1] = 1
				flag3[j + i] = 1
				dfs(append(target, adaptor(j, n)), result)
				flag1[j] = 0
				flag2[i - j + n - 1] = 0
				flag3[j + i] = 0
			}
		}
	}
	dfs([]string{}, &result)

	return result
}

func adaptor(i int, n int) string {
	prefix := strings.Repeat(".", i)
	suffix := strings.Repeat(".", n - i - 1)
	return prefix + "Q" + suffix
}
