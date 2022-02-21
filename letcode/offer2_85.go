package letcode

func generateParenthesis2(n int) []string {
	// 方案1 dfs
	result := make([]string, 0)
	var dfs func(path string, left int, right int)
	dfs = func(path string, left int, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		if left < n {
			dfs(path + "(" , left + 1, right)
		}
		if right < left {
			dfs(path + ")", left, right + 1)
		}

	}
	dfs("", 0 ,0)
	return result
}

func generateParenthesis3(n int) []string {
	// 方案2 dfs
	result := make([]string, 0)
	path := ""
	var backtrack func(left int, right int)
	backtrack = func(left int, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		if left < n {
			path = path + "("
			backtrack(left + 1, right)
			path = path[:len(path) - 1]
		}
		if right < left {
			path = path + ")"
			backtrack(left, right + 1)
			path = path[:len(path) - 1]
		}

	}
	backtrack(0 ,0)
	return result
}

