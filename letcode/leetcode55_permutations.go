package letcode

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	permute1(nums, make([]bool, len(nums)), []int{}, &result)
	return result
}

// 回溯法，记录当前是否被用过，
func permute1(nums []int, used []bool, path []int, result *[][]int) {
	if len(path) == len(nums) {
		*result = append(*result, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			// 递归往下找
			used[i] = true
			permute1(nums, used, append(path, nums[i]), result)
			// 找到了再回来
			used[i] = false
		}
	}
}
