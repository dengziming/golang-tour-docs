package letcode

func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n

	tmp := make([]int, 0)
	tmp = append(tmp, nums[n-k:]...)
	for i := n - 1; i - k >= 0; i-- {
		nums[i] = nums[i - k]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}
