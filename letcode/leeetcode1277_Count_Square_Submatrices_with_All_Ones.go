package letcode

func countSquares(matrix [][]int) int {
	min := func(a,b,c int) int {
		if a < b {
			if a < c {
				return a
			} else {
				return c
			}
		} else {
			if b < c {
				return b
			} else {
				return c
			}
		}
	}

	result := 0

	dp := make([][]int, len(matrix))
	dp[0] = make([]int, len(matrix[0]))
	for i, v := range matrix[0] {
		if v == 1 {
			dp[0][i] = 1
			result += 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		if matrix[i][0] == 1 {
			dp[i][0] = 1
			result += 1
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i - 1][j - 1], dp[i][j - 1], dp[i -1][j]) + 1
				result += dp[i][j]
			}
		}
	}

	return result
}
