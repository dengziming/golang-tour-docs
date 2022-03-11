package letcode

// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
// Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.
// You may assume that you have an infinite number of each kind of coin.
//
// The answer is guaranteed to fit into a signed 32-bit integer.
//
//链接：https://leetcode-cn.com/problems/coin-change-2
//
func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1

	// 这样计算有问题，会重复计算同一个硬币不同顺序
	/*for i := 1; i <= amount; i++ {
		dp[i] = 0
		for _, coin := range coins {
			j := i - coin
			if j >= 0 {
				dp[i] += dp[j]
			}
		}
	}*/

	//
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			j := i - coin
			if j >= 0 {
				dp[i] += dp[j]
			}
		}

	}

	return dp[amount]
}
