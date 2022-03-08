package letcode

import (
	"sort"
)

func minOperations(nums []int) int {
	total := len(nums)
	sort.Ints(nums)

	unique:= func(a []int) []int {
		k := 0
		for _, v := range a[1:] {
			if a[k] != v {
				k++
				a[k] = v
			}
		}
		return a[:k+1]
	}

	nums = unique(nums)
	n := len(nums)

	// dp[i] 代表把最小的改成 nums[i] 需要的操作数
	dp := make([]int, n)
	dp[0] = n - 1
	length := 1
	for i := 1; i < n; i++ {
		l := sort.SearchInts(nums[:i], nums[i]-total+1)
		if i-l+1 > length {
			// [l,r] 内的元素均可以保留
			length = i-l+1
		}
	}
	return total-length
}
