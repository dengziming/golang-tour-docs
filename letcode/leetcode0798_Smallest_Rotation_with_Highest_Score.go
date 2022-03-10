package letcode

// https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/
func bestRotation(nums []int) int {
	// 计算每个 nums[i] +1 所属的下标范围
	// nums[i] +1 ，则 index >= nums[i] 时，可以加分
	// index 一开始为 i. 当 i 大于等于 k，移动后变成 i - k，i 小于 k，移动后变成 i-k+n

	// 当 i < nums[i]，一开始不满足 +1 条件，我们需要让 nums[i] 往右移动
	// if i >= nums[i], 一开始不满足 +1 条件，我们可以适当往左移动，但是也可以往右移动

	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		// low 到 high-1 是对于 num 能加分的 k 的范围
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}
