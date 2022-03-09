package letcode

func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
		return x
	}
		return y
	}

	min := func(x, y int) int {
		if x < y {
		return x
	}
		return y
	}

	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
