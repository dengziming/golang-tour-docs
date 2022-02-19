package letcode

func solveSudoku(board [][]byte)  {
	// 记录每一行每个数字是否出现过
	var row [9][9]bool
	// 记录每一列每个数字是否出现过
	var column [9][9]bool
	// 记录每一个格子每个数字是否出现过
	var block [3][3][9]bool

	// 题目给的数，代表不能动的
	var fixed [9][9]bool

	for i, line := range board {
		for j, v := range line {
			if v != '.' {
				digit := v - '1'
				row[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				fixed[i][j] = true
			} else {

			}
		}
	}

	var dfs func(x int, y int) bool
	// 给点 x,y 填一个数
	dfs = func(x int, y int) bool {
		if x == 9 {
			return true
		}
		nextY := (y + 1) % 9
		nextX := x + (y + 1) / 9
		if fixed[x][y] {
			return dfs(nextX, nextY)
		}
		for i := byte(0); i < 9; i++ {
			if !row[x][i] && !column[y][i] && !block[x/3][y/3][i] {
				row[x][i],column[y][i],block[x/3][y/3][i] = true, true, true
				board[x][y] = i + '1'
				if dfs(nextX, nextY) {
					return true
				}
				row[x][i],column[y][i],block[x/3][y/3][i] = false, false, false
			}
		}
		return false
	}
	dfs(0, 0)
}
