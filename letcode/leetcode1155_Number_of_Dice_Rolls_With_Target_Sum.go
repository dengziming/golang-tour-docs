package letcode

// https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/
// 1155. Number of Dice Rolls With Target Sum
func numRollsToTarget(n int, m int, target int) int {
	// 思路：动态规划
	// 考虑前 x 个骰子能够组合出 y 的方案数

	// dp[i] 表示能组合出 i 方案数
	dp := make([]int, target+1)

	mod := int((1e9)+7)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := target; j >= 0; j-- {
			dp[j] = 0
			for k := 1; k <= m; k++ {
				if j-k >= 0 {
					dp[j] = (dp[j] + dp[j-k]) % mod
				}
			}
		}
	}
	return dp[target]
}
