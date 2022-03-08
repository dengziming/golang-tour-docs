package letcode

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	longest := 1
	result := s[0:1]

	palin := make([][]bool, len(s))
	// 当 j < i 时设置为 true
	for i := 0; i < len(s); i++ {
		palin[i] = make([]bool, len(s))
		for j := 0; j <= i; j++ {
			palin[i][j] = true
		}
	}
	// j > i 时候
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			palin[i][j] = palin[i + 1][j - 1] && s[i] == s[j]
			if palin[i][j] && j + 1 - i > longest {
				longest = j - i
				result = s[i: j + 1]
			}
		}
	}

	return result
}
