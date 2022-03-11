package letcode

func countSubstrings(s string) int {
	n := len(s)
	
	result := 0
	palin := make([][]bool, n)
	for i := 0; i < n; i++ {
		palin[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		palin[i][i] = true
		result = result + 1
		if i < n - 1 && s[i] == s[i + 1] {
			palin[i][i + 1] = true
			palin[i + 1][i] = true
			result = result + 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			palin[i][j] = palin[i + 1][j - 1] && s[i] == s[j]
			if palin[i][j] {
				result = result + 1
			}
		}
	}
	return result
}
