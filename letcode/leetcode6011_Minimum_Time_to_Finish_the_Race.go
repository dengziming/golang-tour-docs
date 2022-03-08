package letcode

import (
	"math"
)

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	// 一个轮胎最多连续使用 n 次，那么 2 的 n-1 次方 不能大于换轮胎的时间。
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	length := min(18, numLaps+1)
	minTime := make([]int, length)
	for i := range minTime {
		minTime[i] = math.MaxInt32
	}
	for _, tire := range tires {
		f, r := tire[0], tire[1]
		for x, time, sum := 1, f, 0; time <= changeTime+f && x < length; x++ {
			sum += time
			minTime[x] = min(minTime[x], sum)
			time *= r
		}
	}

	dp := make([]int, numLaps+1)
	dp[0]=-changeTime
	for i := 1; i <= numLaps; i++ {
		// dp[i] 可以从 dp[j] 一步过来
		dp[i] = math.MaxInt32
		// j 是上一次的起始地址
		for j := 0; j < i; j++ {
			if i-j < length {
				tmp := dp[j] + minTime[i-j]
				if minTime[i-j] != math.MaxInt32 && tmp < dp[i] {
					dp[i] = tmp
				}
			}
		}
		dp[i] += changeTime
	}
	return dp[numLaps]
}
