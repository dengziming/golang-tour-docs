package letcode

func maxSubArray(nums []int) int {
	result := nums[0]
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = dp + nums[i]
		} else {
			dp = nums[i];
		}
		if dp > result {
			result = dp
		}
	}
	return result
}
