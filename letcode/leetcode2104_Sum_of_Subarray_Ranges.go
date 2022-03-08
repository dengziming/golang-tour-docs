package letcode

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	maxDp := make([][]int, n)
	minDp := make([][]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < n; i++ {
		maxDp[i] = make([]int, n)
		minDp[i] = make([]int, n)
		maxDp[i][i] = nums[i]
		minDp[i][i] = nums[i]
		for j := i+1; j < n; j++ {
			maxDp[i][j] = max(maxDp[i][j-1], nums[j])
			minDp[i][j] = min(minDp[i][j-1], nums[j])
		}
	}

	var result int64
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			result+= int64(maxDp[i][j]-minDp[i][j])
		}
	}

	return result
}
