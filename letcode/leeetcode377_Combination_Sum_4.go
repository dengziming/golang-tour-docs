package letcode


func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _,num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 这个交换了 i 和 j 的位置，得到的是不考虑顺序的答案
func combinationSum4_(nums []int, target int) int {
	n := len(nums)
	// dp[i][j] 代表使用前 i 个元素能得到 j 的数量. dp[0][j] 当前仅当整除。 可以递推得到 dp[i+1][j] = ∑dp[i-1][j- 任意自然数*nums[i+1]]
	// 由于 dp[i+1][j] 只依赖 dp[i] 中小于等于 n 的部分，可以进行状态压缩，但是要从 大到小 开始
	dp := make([]int, target+1)
	for j := 0; j <= target; j++ {
		if j%nums[0] == 0 {
			dp[j] = 1
		}
	}

	// 使用前 i 个数
	for i := 1; i < n; i++ {
		// 组成目标 j
		for j := target; j >= 0; j-- {
			// 使用 i 第 k 次
			tmp := 0
			for k := 0; k * nums[i] <= j; k++ {
				tmp += dp[j-k * nums[i]]
			}
			dp[j] = tmp
		}
	}
	return dp[target]
}
