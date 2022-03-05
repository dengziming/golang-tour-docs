package letcode

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	result := 0
	for _, num := range arr {
		dp[num] = dp[num-difference]+1
		if dp[num] > result {
			result = dp[num]
		}
	}
	return result
}
