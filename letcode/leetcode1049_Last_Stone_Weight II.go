package letcode

// https://leetcode-cn.com/problems/last-stone-weight-ii/
func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, num := range stones {
		sum += num
	}
	n := len(stones)
	m := sum/2+1

	// dp[i][j] 代表选择前 i 块石头，能凑成的 j
	dp := make([]bool, m)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := m-1; j >= 0; j-- {
			if j >= stones[i] {
				dp[j] = dp[j] || dp[j-stones[i]]
			}
		}
	}

	for j := m-1; ; j-- {
		if dp[j] {
			return sum - 2*j
		}
	}

}

// 二维 dp
func lastStoneWeightII_2dim(stones []int) int {
	// 思路：dfs，任意挑一个，再挑选一个。。。一共有 n! 种情况，肯定超时
	// 可以看成是 n 个重量相加减的结果，只需要计算 sum-2*neg 大于 0 的最小值

	sum := 0
	for _, num := range stones {
		sum += num
	}
	n := len(stones)
	m := sum/2+1

	// dp[i][j] 代表选择前 i 块石头，能凑成的 j
	dp := make([][]bool, n+1)
	dp[0] = make([]bool, m)
	dp[0][0] = true

	for i := 0; i < n; i++ {
		dp[i+1] = make([]bool, m)
		dp[i+1][0] = true
		for j := 1; j < m; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= stones[i] {
				dp[i+1][j] = dp[i][j] || dp[i][j-stones[i]]
			}
		}
	}

	for j := m-1; ; j-- {
		if dp[n][j] {
			return sum - 2*j
		}
	}
}
