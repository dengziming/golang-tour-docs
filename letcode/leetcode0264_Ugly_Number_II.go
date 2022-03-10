package letcode

func nthUglyNumber(n int) int {
	dp := make([]int, 1)

	// 初识化为 1 开始
	dp[0] = 1

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	p2, p3, p5 := 0, 0, 0
	i := 1
	for ; i <= n; i++ {
		v2 := dp[p2] * 2
		v3 := dp[p3] * 3
		v5 := dp[p5] * 5

		minV := min(v2, min(v3, v5))
		dp = append(dp, minV)
		if v2 == minV {
			p2++
		}
		if v3 == minV {
			p3++
		}
		if v5 == minV {
			p5++
		}
	}

	return dp[n - 1]
}
