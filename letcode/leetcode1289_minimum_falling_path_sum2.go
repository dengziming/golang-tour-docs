package letcode

import "math"

func minFallingPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = grid[i][j]
			} else {
				dp[i][j] = math.MaxInt
				for k := 0; k < n; k++ {
					if k != j && dp[i - 1][k] + grid[i][j] < dp[i][j] {
						dp[i][j] = dp[i - 1][k] + grid[i][j]
					}
				}
			}
		}
	}

	result := dp[m - 1][0]

	for i := 1; i < n; i++ {
		if dp[m - 1][i] < result {
			result = dp[m - 1][i]
		}
	}
	return result
}
