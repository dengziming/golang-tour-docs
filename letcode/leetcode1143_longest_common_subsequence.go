package letcode


func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1) + 1)

	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2) + 1)
		dp[i][0] = 0
		for j := 1; j <= len(text2); j++ {
			if i == 0 {
				dp[i][j] = 0
			} else if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else if dp[i][j - 1] > dp[i - 1][j] {
				dp[i][j] = dp[i][j - 1]
			} else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
