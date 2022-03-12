package letcode

import (
	"math"
)

func findSubsequences(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int,0)

	var dfs func(i int, last int, target []int)
	dfs = func(i int, last int, target []int) {
		if i == n {
			tmp := append([]int{}, target...)
			if len(tmp) > 1 {
				ans = append(ans, tmp)
			}
			return
		}

		// 每个有选或者不选两种情况
		if nums[i] >= last {
			dfs(i + 1, nums[i], append(target, nums[i]))
		}
		if nums[i] != last {
			tmp := append([]int{}, target...)
			dfs(i+1, last, tmp)
		}
	}

	dfs(0, math.MinInt, []int{})
	return ans
}
