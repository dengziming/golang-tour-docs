package letcode

func numberOfGoodSubsets(nums []int) int {
	// 思路: nums 一共就 30 种数，而且答案的每个数只能出现一次，所以可以暴力得到所有的情况，然后检查这个情况是否包含在 nums 中
	// 反过来思考，这是一个背包问题，从 nums 中抽取一定的数，满足一定的条件，直接用动态规划。
	// 动态规划，f[i][j] 代表只使用前 i 个数，使得质因数使用情况为 j 的方案数: f[i][j] = ∑f[i-1][j | nums[j]]

	mod := 1000000007
	// 10 个素数
	primes := []int{2,3,5,7,11,13,17,19,23,29}
	sub := func(value int) int {
		subset := 0
		for i,prime := range primes {
			if value % prime == 0 {
				subset |= 1 << i
			}
		}
		return subset
	}
	// 只有这些数，才能有资格被选中
	targets := []int{2,3,5,6,7,10,11,13,14,15,17,19,21,22,23,26,29,30}

	// 每个数有多少个
	counts := make([]int, 31)
	for _, num := range nums {
		counts[num]++
	}

	// dp[i] 代表以 组合情况为 i，由 dp[len][i] 简化而来
	dp := make([]int, 1 << len(primes))
	dp[0] = 1

	for i := 0; i < counts[1]; i++ {
		dp[0] = (dp[0] * 2) % mod
	}

	for i := 0; i < len(targets); i++ {
		target := targets[i]
		// 没有这个数，直接返回
		if counts[target] == 0 {
			continue
		}

		subset := sub(target)
		for mask := subset; mask <= len(dp) - 1; mask++ {
			// mask 至少要包含了 subset
			if (subset & mask) == subset && dp[mask ^ subset] > 0 {
				dp[mask] = (dp[mask] + dp[mask ^ subset] * counts[target]) % mod
			}
			/*if subset & mask == 0 {
				dp[subset | mask] = (dp[subset | mask] + dp[mask] * counts[target]) % mod
			}*/
		}
	}

	ans := 0
	for i := range dp {
		if i > 0 {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}
