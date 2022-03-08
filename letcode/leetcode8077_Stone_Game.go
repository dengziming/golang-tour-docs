package letcode

func stoneGame(piles []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(piles)
	// dp[i][j] 表示从 [i, j] 中拿能达到的最大差值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		// 从一个拿能达到的最大差值就是 piles[i]
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			// 此时，可以选择 i 和 j
			dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
		}
	}

	return dp[0][n - 1] > 0
}


func stoneGame_2(piles []int) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(piles)
	// dp[i][j] 表示从 [i, j] 中拿能达到的最大差值
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		// 从一个拿能达到的最大差值就是 piles[i]
		dp[i] = piles[i]
	}

	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			// 此时，可以选择 i 和 j
			dp[j] = max(piles[i] - dp[j], piles[j] - dp[j-1])
		}
	}

	return dp[n - 1] > 0
}
