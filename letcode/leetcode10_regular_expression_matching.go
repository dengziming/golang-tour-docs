package letcode

func isMatch(s string, p string) bool {
	// 递归机
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(p) == 0 {
		return false
	}

	// p 一定不为空的情况

	// 还剩下至少两个字符，并且是 .*
	if len(p) > 1 && p[1] == '*' && p[0] == '.'{
		// 缩小正则、缩小字符串
		return isMatch(s, p[2:]) || (len(s) > 0 && isMatch(s[1:], p))
	} else if len(p) > 1 && p[1] == '*' {
		// 还剩下至少两个字符，并且是 x*
		return isMatch(s, p[2:]) || (len(s) > 0 && s[0] == p[0] && isMatch(s[1:], p))
	} else if len(s) == 0 {
		return false
	} else if p[0] == '.' {
		// 接下来是 .
		return isMatch(s[1:], p[1:])
	} else {
		return p[0] == s[0] && isMatch(s[1:], p[1:])
	}
}
