package letcode

import "math"

func numSquares(n int) int {
	count := int(math.Sqrt(float64(n)))

	dp := make([][]int, count)

	IN_VALID := -1

	// 遍历物品
	for i := 0; i < count; i++ {
		num := (i + 1) * (i + 1)
		dp[i] = make([]int, n + 1)
		dp[i][0] = 0
		for j := 0; j <= n; j++ {
			//  dp 的初识状态
			if j == 0 {
				dp[i][0] = 0
			} else if i == 0 {
				dp[0][j] = j
			} else {
				dp[i][j] = IN_VALID
				// 每个数字可以取无数遍，所以遍历
				for k := 0; k * num <= j ; k++ {
					if dp[i-1][j-k*num] != IN_VALID {
						if dp[i][j] == IN_VALID || dp[i-1][j-k*num] + k < dp[i][j] {
							dp[i][j] = dp[i-1][j-k*num] + k
						}
					}

				}
			}
		}
	}
	return dp[count - 1][n]
}
