package letcode

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/super-ugly-number/
// 313. Super Ugly Number
func nthSuperUglyNumber(n int, primes []int) int {
	// 动态规划，计算下一个以 primes[i] 结尾的数即可
	sort.Ints(primes)
	dp := make([]int, n)
	dp[0] = 1

	// 以 primes[i] 结尾的下一个丑逼在 dp 中的索引
	next := make([]int, len(primes))

	for i := 1; i < n; i++ {
		// 从 next 里面选择一个
		min := math.MaxInt
		minIdx := make(map[int]bool)
		for j := range next {
			tmp := dp[next[j]] * primes[j]
			if tmp < min {
				minIdx = make(map[int]bool)
				minIdx[j] = true
				min = tmp
			} else if tmp == min {
				minIdx[j] = true
			}
		}
		dp[i] = min
		for idx := range minIdx {
			next[idx]++
		}
	}
	return dp[n-1]
}
