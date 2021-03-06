package letcode

type NumMatrix struct {
	sum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	// 求一个正方形，先得到 mat 的所有前缀和
	// 先得到前缀和，注意 dp 代表的是不包含 i j 的矩形
	dp := make([][]int, m + 1)
	dp[0] = make([]int, n + 1)
	for i := 1; i < m+1; i++ {
		dp[i] = make([]int, n + 1)
		for j := 1; j < n+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sum: dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
