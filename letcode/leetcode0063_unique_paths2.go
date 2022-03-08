package letcode

// 专题: 动态规划 路径问题
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))

	for i, row := range obstacleGrid {
		dp[i] = make([]int, len(row))
		for j, data := range row {
			if data == 1 {
				dp[i][j] = 0
			} else if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j - 1]
			} else if j == 0 {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
			}
		}
	}
	return dp[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]
}
