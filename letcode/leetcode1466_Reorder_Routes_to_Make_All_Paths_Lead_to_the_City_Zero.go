package letcode

// https://leetcode-cn.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
// 1466. Reorder Routes to Make All Paths Lead to the City Zero
func minReorder(n int, connections [][]int) int {
	// 整体是一棵树，从 0 往外遍历，如果不朝向 0 就 reverse

	// 保存每个点可以去哪里
	edges1 := make([][]int, n)

	// 保存每个点可以从哪里来
	edges2 := make([][]int, n)

	for _, conn := range connections {
		from, to := conn[0], conn[1]
		edges1[from] = append(edges1[from], to)
		edges2[to] = append(edges2[to], from)
	}

	// 检查每个节点
	checked := make([]bool, n)
	var dfs func(i int)
	result := 0
	dfs = func(i int) {
		checked[i] = true
		for _, edge := range edges1[i] {
			if !checked[edge] {
				result++
				dfs(edge)
			}
		}
		for _, edge := range edges2[i] {
			if !checked[edge] {
				dfs(edge)
			}
		}
	}
	dfs(0)
	return result
}
