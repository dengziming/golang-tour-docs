package letcode

func subsetsWithDup(nums []int) [][]int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	var dfs func(i int, sub []int)

	result := make([][]int, 0)
	dfs = func(i int, sub []int) {
		result = append(result, append([]int{}, sub...))
		if i == len(nums) {
			return
		}
		for num := range counts {
			if counts[num] > 0 && (len(sub) == 0 || num >= sub[len(sub)-1]) {
				counts[num]--
				dfs(i+1, append(sub, num))
				counts[num]++
			}
		}
	}
	dfs(0, []int{})
	return result
}
