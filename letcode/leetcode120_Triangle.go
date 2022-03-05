package letcode

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	dp[0] = []int{triangle[0][0]}

	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, i + 1)
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		for j := 1; j < i; j++ {
			if dp[i-1][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	min := dp[len(triangle) - 1][0]
	for i := 1; i < len(triangle); i++ {
		if dp[len(triangle) - 1][i] < min {
			min = dp[len(triangle) - 1][i]
		}
	}
	return min
}
