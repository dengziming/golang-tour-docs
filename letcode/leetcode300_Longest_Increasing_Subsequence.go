package letcode

// https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-by-leetcode-soluti/
// 思路1： 动态规划
// 思路2: 贪心：要最长增长子串，就是要让增长速度。可以记录每个元素当前所属排名
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	greedy := []int{nums[0]}
	length := 1

	for i := 1; i < n; i++ {
		num := nums[i]
		if num > greedy[length-1] {
			length++
			greedy = append(greedy, num)
		} else if num <= greedy[0]{
			greedy[0] = num
		} else {
			left, right := 0, len(greedy) - 1
			// 二分查找第一个小于 num 的元素
			// [0 1], 1 =>
			for left < right {
				mid := left + (right - left + 1)/2
				if greedy[mid] >= num {
					right = mid-1
				} else {
					left = mid
				}
			}
			if num < greedy[left+1] {
				greedy[left+1] = num
			}
		}
	}
	return length
}
