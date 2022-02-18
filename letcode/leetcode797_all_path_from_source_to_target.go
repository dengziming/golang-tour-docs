package letcode

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	n := len(graph)

	var dfs func(index int, target []int)
	dfs = func(index int, target []int) {
		if index == n - 1 {
			result = append(result, append(append([]int(nil), target...), n - 1))
			return
		}
		// 下一跳
		for _, next := range graph[index] {
			dfs(next, append(target, index))
		}
	}
	dfs(0, []int{})
	return result
}
