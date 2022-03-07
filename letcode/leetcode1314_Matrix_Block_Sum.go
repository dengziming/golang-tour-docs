package letcode

func matrixBlockSum(mat [][]int, k int) [][]int {
	m := len(mat)
	n := len(mat[0])
	// 求一个正方形，先得到 mat 的所有前缀和
	// 先得到前缀和，注意 dp 代表的是不包含 i j 的矩形
	dp := make([][]int, m + 1)
	dp[0] = make([]int, n + 1)
	for i := 1; i < m+1; i++ {
		dp[i] = make([]int, n + 1)
		for j := 1; j < n+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1] + mat[i-1][j-1]
		}
	}

	// i j 到 k l
	sum := func(i, j, k, l int) int {
		if i < 0 {
			i = 0
		}
		if j < 0 {
			j = 0
		}
		if k > m {
			k = m
		}
		if l > n {
			l = n
		}
		return dp[k][l] - dp[k][j] - dp[i][l] + dp[i][j]
	}

	answer := make([][]int, m)
	for i := 0; i < m; i++ {
		answer[i] = make([]int, n)
		for j := 0; j < n; j++ {
			answer[i][j] = sum(i-k,j-k,i+k+1,j+k+1)
		}
	}
	return answer
}
