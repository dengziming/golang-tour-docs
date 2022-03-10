package letcode

func maximalSquare(matrix [][]byte) int {
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
	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 以 i j 为终点的最大正方形
	dp := make([][]int, len(matrix))
	maxV := 0
	dp[0] = make([]int, len(matrix[0]))
	for i, b := range matrix[0] {
		if b == '1' {
			dp[0][i] = 1
			maxV = max(1, maxV)
		}
	}

	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxV = max(1, maxV)
		}

		for j := 1; j < len(matrix[i]); j++ {
			b := matrix[i][j]
			if b == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				maxV = max(dp[i][j], maxV)
			}
		}
	}
	return maxV * maxV
}
