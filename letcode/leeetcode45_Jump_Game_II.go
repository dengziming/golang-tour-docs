package letcode

func jump(nums []int) int {
	n := len(nums)
	end := n-1
	jump := 0

	for end > 0 {
		for i := 0; i < end; i++ {
			if i + nums[i] >= end {
				end = i
				jump++
				break
			}
		}
	}
	return jump
}
