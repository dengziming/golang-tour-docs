package letcode

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
// 链接：https://leetcode-cn.com/problems/word-break
func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	// 递归版本，计算超时了
	/*var dfs func(s string) bool
	dfs = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		for _, word := range wordDict {
			if len(word) <= len(s) && s[:len(word)] == word && dfs(s[len(word):]) {
				return true
			}
		}
		return false
	}*/

	// dp 版本
	result := make([]bool, n)
	for i := 0; i < n; i++ {
		result[i] = false
		for j := 0; j <= i; j++ {
			for _, word := range wordDict {
				//
				if (j == 0 || result[j - 1]) && len(word) == (i + 1 - j) && word == s[j: i + 1] {
					result[i] = true
					break
				}
			}
		}
	}

	return result[n - 1]

	// return dfs(s)
}
