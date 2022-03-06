package letcode

import (
	"math"
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	// 找到一个 i，使得 小于 i 的所有自然数，有 k 个不在 nums 中。

	limit := len(nums) + k
	sort.Ints(nums)
	nums = append(nums, math.MaxInt)

	var sum = 0
	counts := make(map[int]bool)
	for i := range nums {
		// 重复的直接忽略
		if !counts[nums[i]] {

			// len(counts) 个数 在 nums 中，如果满足这个条件，说明已经有超过 k 数了
			if nums[i] - len(counts) > k {
				// len(counts) 代表已经有的所有元素个数，加上 k 代表填充以后的数组的最大
				limit = k + len(counts)
				break
			}
			sum += nums[i]
			counts[nums[i]] = true
		}
	}

	var result = int64(limit+1)*int64(limit)/2
	return result - int64(sum)

}
