package letcode

func removeInvalidParentheses(s string) []string {
	// 得到需要删除的左右括号个数
	leftToDelete, rightToDelete := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftToDelete ++
		} else if s[i] == ')' {
			if leftToDelete > 0 {
				leftToDelete--
			} else {
				rightToDelete++
			}
		}
	}

	result := make(map[string]bool)
	tmp := 0

	var backtrack func(path string, i int)
	backtrack = func(path string, i int) {
		// 成功
		if i == len(s) && leftToDelete == 0 && rightToDelete == 0 {
			result[path] = true
			return
		}
		// 失败
		if i == len(s) {
			return
		}

		if s[i] == '(' {
			tmp++
			backtrack(path + string(s[i]), i + 1)
			tmp--
			if leftToDelete > 0 {
				leftToDelete--
				backtrack(path, i + 1)
				leftToDelete++
			}
		} else if s[i] == ')' {
			if tmp > 0 {
				tmp--
				backtrack(path + string(s[i]), i + 1)
				tmp++
			}
			if rightToDelete > 0 {
				rightToDelete--
				backtrack(path, i + 1)
				rightToDelete++
			}
		} else {
			backtrack(path + string(s[i]), i + 1)
		}
	}

	backtrack("", 0)

	ans := make([]string, 0)
	for k := range result {
		ans = append(ans, k)
	}
	return ans
}

func removeInvalidParentheses2(s string) []string {
	// 得到需要删除的左右括号个数
	leftToDelete, rightToDelete := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftToDelete ++
		} else if s[i] == ')' {
			if leftToDelete > 0 {
				leftToDelete--
			} else {
				rightToDelete++
			}
		}
	}

	result := make(map[string]bool)

	var dfs func(path string, i, leftToDelete, rightToDelete, tmp int)
	dfs = func(path string, i, leftToDelete, rightToDelete, tmp int) {
		// 成功
		if i == len(s) && leftToDelete == 0 && rightToDelete == 0 {
			result[path] = true
			return
		}
		// 失败
		if i == len(s) {
			return
		}

		if s[i] == '(' {
			dfs(path + string(s[i]), i + 1, leftToDelete, rightToDelete, tmp+1)
			if leftToDelete > 0 {
				dfs(path, i + 1, leftToDelete - 1, rightToDelete, tmp)
			}
		} else if s[i] == ')' {
			if tmp > 0 {
				dfs(path + string(s[i]), i + 1, leftToDelete, rightToDelete, tmp-1)
			}
			if rightToDelete > 0 {
				dfs(path, i + 1, leftToDelete, rightToDelete - 1, tmp)
			}
		} else {
			dfs(path + string(s[i]), i + 1, leftToDelete, rightToDelete, tmp)
		}
	}

	dfs("", 0, leftToDelete, rightToDelete, 0)

	ans := make([]string, 0)
	for k := range result {
		ans = append(ans, k)
	}
	return ans
}
