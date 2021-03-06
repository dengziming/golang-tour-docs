package letcode
func shortestPalindrome(s string) string {
	n := len(s)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i - 1]
		for j != -1 && s[j + 1] != s[i] {
			j = fail[j]
		}
		if s[j + 1] == s[i] {
			fail[i] = j + 1
		}
	}
	best := -1
	for i := n - 1; i >= 0; i-- {
		for best != -1 && s[best + 1] != s[i] {
			best = fail[best]
		}
		if s[best + 1] == s[i] {
			best++
		}
	}
	add := ""
	if best != n - 1 {
		add = s[best + 1:]
	}
	b := []byte(add)
	for i := 0; i < len(b) / 2; i++ {
		b[i], b[len(b) - 1 -i] = b[len(b) - 1 -i], b[i]
	}
	return string(b) + s
}

// 二维数组内存溢出
func shortestPalindrome2(s string) string {
	n := len(s)
	dp := make([][]bool, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i > 0 {
			dp[i][i-1] = true
		}
		for j := i+1; j < n; j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}

	idx := 0
	for i := n - 1; i >= 0; i-- {
		if dp[0][i] {
			idx = i
			break
		}
	}
	result := ""
	for i := n -1; i > idx; i-- {
		result = result+string(s[i])
	}
	return result + s
}
