package letcode

// 一次扫描得到数组每个数左边的最大值
// 一次扫描得到数组每个数右边的最大值
// 每个数左边的最大值和右边的最大值的最小值，减去这个数
func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	max := 0
	left[0] = 0
	for i := 1; i < len(height); i++ {
		if height[i - 1] > max {
			max = height[i - 1]
		}
		left[i] = max
	}

	max = 0
	right[len(height) - 1] = 0
	for i := len(height) - 2; i > -1; i-- {
		if height[i + 1] > max {
			max = height[i + 1]
		}
		right[i] = max
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		var proof int
		if left[i] > right[i] {
			proof = right[i]
		} else {
			proof = left[i]
		}
		if proof > height[i] {
			sum += proof - height[i]
		}
	}
	return sum
}
