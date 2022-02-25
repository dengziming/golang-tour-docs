package letcode

func maxCoins(nums []int) int {
	// 前后各加一个 1
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	// dp[i][j] 代表以 i 和 j 为边界戳破里面的可以获得的最大受益
	// 那么
	dp := make([][]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(nums))
		if i < len(nums)-1 {
			dp[i][i+1] = 0
		}
		for j := i + 2; j < len(nums); j++ {
			maxV := 0
			for k := i+1; k < j; k++ {
				tmp := dp[i][k] + dp[k][j] + nums[i] * nums[k] * nums[j]
				if tmp > maxV {
					maxV = tmp
				}
			}
			dp[i][j] = maxV
		}
	}

	return dp[0][len(nums) - 1]
}

// 递归+cache
func maxCoins2(nums []int) int {
	// 前后各加一个1
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	cache := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]int, len(nums))
	}

	// 最简单分而治之，代表如果只有 i j 两个元素，再戳破一个应该戳破哪个
	var divide func(left, right int) int
	divide = func(left, right int) int {
		if left >= right - 1 {
			return 0
		}
		if cache[left][right] != 0 {
			return cache[left][right]
		}
		maxV := 0
		for i := left + 1; i < right; i++ {
			tmp := divide(left, i) + divide(i, right) + nums[left] * nums[right] * nums[i]
			if tmp > maxV {
				maxV = tmp
			}
		}
		cache[left][right] = maxV
		return maxV
	}

	return divide(0, len(nums) - 1)
}
