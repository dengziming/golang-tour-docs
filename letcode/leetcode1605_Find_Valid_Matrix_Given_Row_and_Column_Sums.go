package letcode

// https://leetcode-cn.com/problems/find-valid-matrix-given-row-and-column-sums/
// 1605. Find Valid Matrix Given Row and Column Sums
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	// 贪心

	m, n := len(rowSum), len(colSum)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	for i,j := 0,0; i < m && j < n; {
		// 选列
		if rowSum[i] > colSum[j] {
			result[i][j] = colSum[j]
			rowSum[i] -= colSum[j]
			j++
		} else { // 选行
			result[i][j] = rowSum[i]
			colSum[j] -= rowSum[i]
			i++
		}
	}
	return result
}


// 超时了
func restoreMatrix_dfs(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	var dfs func(i,j int) bool

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	dfs = func(row, col int) bool {
		// 每次行末需要判断
		if col == 0 && row-1 >= 0 && rowSum[row-1] != 0 {
			return false
		}
		// 最后一行的时候，每一列都必须为 0
		if row == m-1 && col-1 >= 0 && colSum[col-1] != 0{
			return false
		}
		if row == m {
			// 验证一下最后一列
			if colSum[n-1] != 0 {
				return false
			}
			return true
		}
		for i := 0; i <= rowSum[row] && i <= colSum[col]; i++ {
			rowSum[row] -= i
			colSum[col] -= i
			result[row][col] = i

			tmp := dfs(row + (col + 1) / n, (col + 1) % n)
			colSum[col] += i
			rowSum[row] += i
			if tmp {
				return tmp
			}
		}
		return false
	}
	dfs(0, 0)
	return result
}
