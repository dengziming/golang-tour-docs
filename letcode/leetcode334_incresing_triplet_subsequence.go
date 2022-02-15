package letcode

func increasingTriplet(nums []int) bool {
	// 超时了
	/*dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] + 1 > dp[j] {
				dp[i] = dp[j] + 1
				if dp[i] == 3 {
					return true
				}
			}
		}
	}
	return false*/

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		// 二分查找 dp[0, max) 得到 num
		index := binarySearch(&dp, nums[i], 0, maxLen - 1)

		if index == maxLen {
			maxLen ++
			dp[index] = nums[i]
		}
		if dp[index] > nums[i] {
			dp[index] = nums[i]
		}
		if maxLen >= 3 {
			return true
		}
	}
	return false
}

// [1, 3, 5], 2  => 1
// [1, 3, 5], 7  => 3
// [1, 3, 5], 0  => 0
func binarySearch(dp *[]int, value int, left int, right int) int {
	if left == right {
		if (*dp)[left] >= value {
			return left
		}
		return left + 1
	}
	middle := left + right / 2
	if (*dp)[middle] == value {
		return middle
	} else if (*dp)[middle] > value {
		return binarySearch(dp, value, left, middle)
	} else {
		return binarySearch(dp, value, middle + 1, right)
	}
}
