package letcode

func totalNQueens(n int) int {
	result := 0

	var dfs func(target []int)

	var flag1 = make([]int, n)
	var flag2 = make([]int, 2 * n - 1)
	var flag3 = make([]int, 2 * n - 1)

	dfs = func(target []int) {
		if len(target) == n {
			result ++
			return
		}
		i := len(target)
		for j := 0; j < n ; j++ {
			if flag1[j] == 0 && flag2[i - j + n -1] == 0 && flag3[i + j] == 0 {
				flag1[j] = 1
				flag2[i - j + n -1] = 1
				flag3[i + j] = 1
				dfs(append(target, j))
				flag1[j] = 0
				flag2[i - j + n -1] = 0
				flag3[i + j] = 0
			}
		}
	}

	dfs([]int{})
	return result
}
