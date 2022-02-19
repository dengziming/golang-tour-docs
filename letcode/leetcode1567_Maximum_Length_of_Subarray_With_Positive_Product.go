package letcode

// 找到所有负数
func getMaxLen(nums []int) int {

	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	getNegIndies := func(nums []int) []int {
		result := make([]int, 0)
		for i, num := range nums {
			if num < 0 {
				result = append(result, i)
			}
		}
		return result
	}

	left := 0
	maxLen := 0

	n := len(nums)
	for right := 0; right <= n ; right++ {
		if right == n || nums[right] == 0 {
			indies := getNegIndies(nums[left: right])
			subLen := len(indies)
			if subLen % 2 == 0 {
				maxLen = max(right - left, maxLen)
			} else {
				tmp1 := right - left - (indies[0] + 1)
				tmp2 := indies[subLen - 1]
				maxLen = max(maxLen, max(tmp1, tmp2))
			}
			left = right + 1
		}
	}

	return maxLen
}
