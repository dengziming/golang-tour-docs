package letcode

func longestSubstring(s string, k int) int {
	
	var dfs func(s string) int
	
	dfs = func(s string) int {
		if len(s) == 0 {
			return 0
		}
		counts := make(map[uint8]int)
		for i := range s {
			counts[s[i]] = counts[s[i]] + 1
		}

		result := 0
		start := 0
		end := 0
		flag := true
		for i := range s {
			if counts[s[i]] < k {
				flag = false
				end = i
				tmp := dfs(s[start: end])
				if tmp > result {
					result = tmp
				}
				start = end + 1
			}
		}
		if flag {
			return len(s)
		} else {
			end = len(s)
			tmp := dfs(s[start: end])
			if tmp > result {
				result = tmp
			}
			return result
		}
	}
	return dfs(s)
}
