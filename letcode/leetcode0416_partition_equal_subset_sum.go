package letcode

func canPartition(nums []int) bool {
	var target int
	for _, num := range nums {
		target += num
	}
	if target & 1 == 1 {
		return false
	}
	target /= 2

	// 从 nums 中找到相加为 target 的数

	// dp[subTarget][j] 代表前 j 个数字是否能组成 subTarget
	dp := make([][]bool, target + 1)
	for subTarget := 0; subTarget <= target; subTarget++ {
		dp[subTarget] = make([]bool, len(nums))

		for j := 0; j < len(nums); j++ {
			if subTarget == 0 {
				dp[subTarget][j] = true
			} else if j == 0 {
				if subTarget == nums[j] {
					dp[subTarget][j] = true
				} else {
					dp[subTarget][j] = false
				}
			} else {
				dp[subTarget][j] = dp[subTarget][j - 1]

				if subTarget >= nums[j] {
					dp[subTarget][j] = dp[subTarget][j] || dp[subTarget- nums[j]][j - 1]
				}

			}
		}
	}

	return dp[target][len(nums) - 1]
}
