package letcode

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	bit := 0
	for i := 0; i < n; i++ {
		bit = bit | nums[i]
	}
	var dfs func(i int, target int)

	result := 0
	dfs = func(i int, target int) {
		if i == n {
			if target == bit {
				result++
			}
			return
		}

		dfs(i+1, target | nums[i])
		dfs(i+1, target)
	}
	dfs(0, 0)
	return result
}
