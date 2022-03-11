package letcode

func updateBoard(board [][]byte, click []int) [][]byte {
	m := len(board)
	n := len(board[0])
	var dfs func(i, j int)

	inboard := func(i, j int) bool {
		if i >=0 && j >= 0 && i < m && j < n {
			return true
		}
		return false
	}

	dfs = func(i, j int) {
		if board[i][j] == 'M' {
			board[i][j] = 'X'
			return
		}
		if board[i][j] == 'E' {
			mineCnt := 0
			for row := i-1; row <= i+1; row++ {
				for col := j-1; col <= j+1; col++ {
					if inboard(row, col) && board[row][col] == 'M' {
						mineCnt++
					}
				}
			}
			if mineCnt == 0 {
				board[i][j] = 'B'
				for row := i-1; row <= i+1; row++ {
					for col := j-1; col <= j+1; col++ {
						if inboard(row, col) {
							dfs(row, col)
						}
					}
				}
			} else {
				board[i][j] = '0' + byte(mineCnt)
			}
		}
	}

	dfs(click[0], click[1])
	return board
}
