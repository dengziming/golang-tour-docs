package letcode

func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	cache := make([][]bool, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]bool, n)
	}

	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !cache[i][j] && board[i][j] == 'X'{
				for k := j + 1; k < n && board[i][k] == 'X'; k++ {
					cache[i][k] = true
				}
				for k := i + 1; k < m && board[k][j] == 'X'; k++ {
					cache[k][j] = true
				}
				cnt++
			}
		}
	}
	return cnt
}
