package letcode

import "math"

// 专题: 动态规划 路径问题
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	cache := make([][]int, len(grid))
	for i, row := range grid {
		cache[i] = make([]int, len(row))
		for j := range row {
			cache[i][j] = -1
		}
	}

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if cache[i][j] != -1 {
			return cache[i][j]
		}

		var v int
		if i == 0 && j == 0 {
			v = grid[0][0]
		} else if i == 0 {
			v =  grid[i][j] + dfs(i, j - 1)
		} else if j == 0 {
			v =  grid[i][j] + dfs(i - 1, j)
		} else {
			tmp := math.Min(float64(dfs(i-1, j)), float64(dfs(i, j-1)))
			v = int(tmp) + grid[i][j]
		}
		cache[i][j] = v
		return v
	}

	return dfs(len(grid) -1 , len(grid[0]) -1)
}
