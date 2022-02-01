package letcode

import "unicode"

// 思路：找到只出现一次的字符，找不到就是完美的，找到了就按照这些字符分割，递归计算子串
func longestNiceSubstring(s string) string {

	var dfs func(s string) string

	dfs = func(s string) string {
		if len(s) == 0 {
			return ""
		}
		counts := make(map[uint8]int)

		for i := range s {
			ch := s[i]
			if unicode.IsLower(rune(ch)) {
				counts[s[i]] = counts[ch] | 1
			} else {
				ch = uint8(unicode.ToLower(rune(ch)))
				counts[ch] = counts[ch] | 2
			}

		}

		result := ""
		start := 0
		end := 0
		isPerfect := true // 是不是不存在只出现一次的字符
		for i := 0; i < len(s); i++ {
			ch := s[i]
			if unicode.IsUpper(rune(ch)) {
				ch = uint8(unicode.ToLower(rune(ch)))
			}
			if counts[ch] != 3 {
				isPerfect = false
				end = i
				tmp := dfs(s[start: end])
				if len(tmp) > len(result) {
					result = tmp
				}
				start = end + 1
			}
		}

		if isPerfect {
			return s
		} else {
			// 最后还有一节
			end = len(s)
			tmp := dfs(s[start: end])
			if len(tmp) > len(result) {
				result = tmp
			}
		}
		return result
	}
	return dfs(s)
}
