package letcode

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true

	for i := 0; i < n; i++ {
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < n; j++ {
				dp[i+j] = true
			}
		}

	}
	return dp[n-1]
}
