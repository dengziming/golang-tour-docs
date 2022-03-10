package letcode

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) <= 1 {
		return len(nums)
	}

	max := func(a int, b int) int {
		if a > b {
		return a
	}
		return b
	}

	// positive[i] 代表以 nums[i] 结尾的 wiggle seq，并且最近一跳为正
	// negative 类似
	positive := make([]int, n)
	negative := make([]int, n)

	positive[0] = 1
	negative[0] = 1

	maxV := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				continue
			}
			if nums[i] > nums[j] {
				positive[i] = max(positive[i], negative[j]+1)
				maxV = max(maxV, positive[i])
			} else {
				negative[i] = max(negative[i], positive[j]+1)
				maxV = max(maxV, negative[i])
			}
		}
	}

	return maxV
}
