package letcode

// 思路1: 保存一个长度为 k 的堆，堆顶的就是最大值，每次移动窗口就移出一个元素，然后添加一个元素，时间复杂度较高
// 思路2: 由于只需要最大值，复杂的是删除，我们保留当前最大和次大，删除了最大就用次大顶替，但是如果删除了次大呢？用什么顶替？这个思路不通，但是如果能保证删除的一定不是次大呢，类似的有第三种思路，
// 思路3: 由于我们只关心最大的元素，维护一个排序的队列，每次
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)

	// 元素对应的 index 入队列，已有的元素的 index 肯定更小，如果值也跟小那就没有存在的意义
	push := func(a int) {
		for len(queue) > 0 && nums[queue[len(queue) - 1]] <= nums[a] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, a)
	}

	for i := 0; i < k - 1; i++ {
		push(i)
	}

	n := len(nums)
	result := make([]int, 0)
	for i := k - 1; i < n; i++ {
		// 如果当前的头已经在窗口外了，那留着也没有意义了
		if len(queue) > 0 && i - queue[0] >= k {
			queue = queue[1:]
		}
		push(i)
		result = append(result, nums[queue[0]])
	}
	return result
}
