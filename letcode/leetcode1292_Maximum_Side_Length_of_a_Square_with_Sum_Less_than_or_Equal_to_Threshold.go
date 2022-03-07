package letcode

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	// 求一个正方形，先得到 mat 的所有前缀和
	// 先得到前缀和，注意 dp 代表的是不包含 i j 的矩形
	dp := make([][]int, m + 1)
	dp[0] = make([]int, n + 1)
	for i := 1; i < m+1; i++ {
		dp[i] = make([]int, n + 1)
		for j := 1; j < n+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1] + mat[i-1][j-1]
		}
	}

	// 从 i j 开始，长 k 的矩形的和
	sum := func(i, j, k int) int {
		return dp[i+k][j+k] - dp[i][j+k] - dp[i+k][j] + dp[i][j]
	}

	maxLen := 0
	// 遍历每一个前缀和，得到最大的半径
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 遍历每一个正方形
			for k := maxLen+1; i + k <= m && j + k <= n; k++ {
				if sum(i, j, k) <= threshold {
					maxLen = k
				} else {
					break
				}
			}
		}
	}
	return maxLen
}
