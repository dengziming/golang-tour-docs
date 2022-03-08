package letcode

// https://leetcode-cn.com/problems/combination-sum
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	current := []int{}
	var dfs func(target, idx int)
	dfs = func(target int, index int) {
		if target < 0 {
			return
		}
		if target == 0 {
			result = append(result, append([]int(nil), current...))
			return
		}

		for i := index; i < len(candidates); i++ {
			current = append(current, candidates[i])
			dfs(target - candidates[i], i)
			current = current[:len(current) - 1]
		}
	}

	dfs(target, 0)
	return result
}


func combinationSum1(candidates []int, target int, current []int, result *[][]int, index int)  {
	if target == 0 {
		*result = append(*result, current)
		return
	} else if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		combinationSum1(candidates, target - candidates[i], append(current, candidates[i]), result, i)
	}
}
