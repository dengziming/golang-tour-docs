package letcode

func repeatedSubstringPattern(s string) bool {
	return kmp(s + s, s)
}

// 判断一个 pattern 是否在 query 字符串中出现
func kmp(query, pattern string) bool {
	// 先计算 pattern 每个位置匹配失败后可以移动的位数
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i - 1]
		for j != -1 && pattern[j + 1] != pattern[i] {
			// // 上一个字符匹配成功了，但是当前字符匹配失败了，并不能直接设置为 0，而是需要循环往前直到匹配成功
			j = fail[j]
		}
		if pattern[j + 1] == pattern[i] {
			fail[i] = j + 1
		}
	}

	match := -1
	for i := 1; i < n - 1; i++ {
		for match != -1 && pattern[match + 1] != query[i] {
			// 失败了
			match = fail[match]
		}
		if pattern[match + 1] == query[i] {
			match++
			if match == m - 1 {
				return true
			}
		}
	}

	return false
}
