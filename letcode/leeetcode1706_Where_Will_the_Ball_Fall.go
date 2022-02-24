package letcode


func findBall(grid [][]int) []int {
	// 多少层
	n := len(grid)
	// 多少个球
	m := len(grid[0])
	result := make([]int,m)

	for i := 0; i < m; i++ {
		// 初始位置
		col := i
		for j := 0; j < n; j++ {
			dir := grid[j][col]
			col += dir
			if col < 0 || col >= m || grid[j][col] != dir{
				col = -1
				break
			}
		}
		result[i] = col
	}
	return result
}
