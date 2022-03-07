package letcode

func dieSimulator(n int, rollMax []int) int {
	var mod int64 = 1000000007
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, 6)
	}

	dp[0] = []int64{1,1,1,1,1,1}
	// 第 i 次扔
	for i := 1; i < n; i++ {
		for j := 0; j < 6; j++ {
			var tmp int64
			// 每一个都可以基于上一次的结果，加上 6 种选择
			for k := 0; k < 6; k++ {
				tmp = tmp + dp[i-1][k]
			}
			// 6 中选择中，对于本次刚好等于 j 的，需要判断是否超过了 rollMax 的限制
			// 加入 rollMax[j] == 3，需要减去前 3 次都是 j，而往前 4 次不是 j 的
			if i - rollMax[j] > 0 {
				for k := 0; k < 6; k++ {
					if k != j {
						tmp = tmp - dp[i - rollMax[j] - 1][k]
					}
				}
			} else if i - rollMax[j] == 0 {
				// 如果刚好从一开始到现在都是这个数
				tmp = tmp - 1
			}

			dp[i][j] = (tmp + 6*mod) % mod
		}
	}
	var result int64 = 0
	for _, cnt := range dp[n-1] {
		result = (result + cnt) % mod
	}
	return int(result)
}

// dfs 超时
func dieSimulator1(n int, rollMax []int) int {
	result := 0

	// i 当前第几次投
	// pre 上一次是多少
	// 上一个多少次了
	var dfs func(i int, pre int, cnt int)

	dfs = func(i int, pre int, cnt int) {
		if i == n {
			result++
			return
		}
		for j := 0; j < 6; j++ {
			if pre != j {
				dfs(i+1, j, 1)
			} else if cnt < rollMax[j] {
				dfs(i+1, j, cnt + 1)
			}
		}
	}

	dfs(0, -1, 0)
	return result
}
