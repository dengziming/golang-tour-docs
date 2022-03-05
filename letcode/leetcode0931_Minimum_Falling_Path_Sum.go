package letcode

import (
	"math"
)

func minFallingPathSum1(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	dp[0] = matrix[0]

	min := func(a int, b int, c int) int {
		if a < b {
			if a < c {
				return a
			} else {
				return c
			}
		}
		if b < c {
			return b
		}
		return c
	}
	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		row := matrix[i]
		dp[i][0] = min(dp[i-1][0], dp[i-1][1], math.MaxInt) + row[0]
		for j := 1; j < len(row) - 1; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j - 1], dp[i-1][j + 1]) + row[j]
		}
		dp[i][len(row) - 1] = min(dp[i-1][len(row) - 2], dp[i-1][len(row) - 1], math.MaxInt) + row[len(row) - 1]
	}

	minV := dp[len(matrix) - 1][0]
	for i := 1; i < len(matrix[0]); i++ {
		if dp[len(matrix) - 1][i] < minV {
			minV = dp[len(matrix) - 1][i]
		}
	}
	return minV
}
