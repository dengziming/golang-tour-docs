package letcode

import "math"

func maxSubarraySumCircular(nums []int) int {
	// 以当前数结尾的最大的 sum
	maxDp, minDp := 0, 0
	// 所有的 maxDp
	maxV, minV := nums[0], nums[0]

	sum := 0

	// 全是负数
	allNeg := true
	maxNeg := math.MinInt

	for _, num := range nums {
		maxDp = max(maxDp + num, num)
		minDp = min(minDp + num, num)

		maxV = max(maxDp, maxV)
		minV = min(minDp, minV)
		sum += num

		if num > 0 {
			allNeg = false
		} else {
			maxNeg = max(maxNeg, num)
		}
	}

	if allNeg {
		return maxNeg
	}

	return max(maxV, sum - minV)
}
