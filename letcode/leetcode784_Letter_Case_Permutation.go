package letcode

func letterCasePermutation(s string) []string {
	var ans []string

	var dfs func(str []byte, start int)
	dfs = func(str []byte, start int) {
		ans = append(ans, string(str))
		for i:=start; i<len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' {
				str[i] -= 32
				dfs(str, i+1)
				str[i] += 32
			}
			if str[i] >= 'A' && str[i] <= 'Z' {
				str[i] += 32
				dfs(str, i+1)
				str[i] -= 32
			}
		}
	}

	dfs([]byte(s), 0)
	return ans
}
