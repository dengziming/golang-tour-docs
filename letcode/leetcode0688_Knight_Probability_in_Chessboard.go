package letcode

import "math"

func knightProbability(n int, k int, row int, column int) float64 {
	var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}
	inBoard := func(row int, column int) bool {
		return row>=0 && column>=0 && row <=n-1 && column<=n-1
	}
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; inBoard(x, y) {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// 递归
func knightProbability2(n int, k int, row int, column int) float64 {
	cnt := 0
	total := math.Pow(8, float64(k))
	inBoard := func(row int, column int) bool {
		return row>=0 && column>=0 && row <=n-1 && column<=n-1
	}
	var dfs func(i int, row int, column int)

	dfs = func(i int, row int, column int) {
		if !inBoard(row, column) {
			return
		}
		if i == 0 {
			cnt += 1
			return
		}
		/*for j := -2; j < 2; j++ {
			for k := -2; k < 2; k++ {
				if j*k != 0 {

				}
			}
		}*/
		dfs(i-1, row+1, column+2)
		dfs(i-1, row+2, column+1)
		dfs(i-1, row+2, column-1)
		dfs(i-1, row+1, column-2)
		dfs(i-1, row-1, column-2)
		dfs(i-1, row-2, column-1)
		dfs(i-1, row-2, column+1)
		dfs(i-1, row-1, column+2)
	}

	dfs(k, row, column)
	return float64(cnt)/total
}
