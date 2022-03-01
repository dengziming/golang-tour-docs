package letcode

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxV := 1
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		if i != 0 {
			// 设置为 0 代表这个回文串是空的
			dp[i][i-1] = 0
		}
	}
	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			// 这里改一下
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1] + 2)
			}
			dp[i][j] = max(dp[i][j], dp[i+1][j])
			dp[i][j] = max(dp[i][j], dp[i][j-1])

			if dp[i][j] > maxV {
				maxV = dp[i][j]
			}
		}
	}
	fmt.Printf("%v", dp)
	return maxV
}
