package letcode

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
//
// https://leetcode-cn.com/problems/permutations-ii
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)

	// 记录每个数字出现的个数
	counts := make(map[int]int)
	// 去重后的数字
	set := make(map[int]bool)
	for _, num := range nums {
		if _, ok := counts[num]; !ok {
			counts[num] = 0
		}
		counts[num] = counts[num] + 1
		set[num] = true
	}

	var dfs func(set map[int]bool, target []int, result *[][]int)
	dfs = func(set map[int]bool, target []int, result *[][]int) {
		if len(target) == len(nums) {
			// TODO 这个的目的是啥，为啥 *result = append(*result, target) 就会指针错乱
			*result = append(*result, append([]int(nil), target...))
			return
		}
		for num := range set {
			if counts[num] > 0 {
				counts[num] = counts[num] - 1
				dfs(set, append(target, num), result)
				counts[num] = counts[num] + 1
			}
		}
	}
	
	dfs(set, []int{}, &result)
	return result
}
