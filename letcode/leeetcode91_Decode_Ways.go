package letcode

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 1
	}

	dp1, dp2 := 0, 0
	dp1 = 1

	if s[0:1] == "0" {
		dp2 = 0
	} else {
		dp2 = 1
	}
	if n ==1 {
		return dp2
	}

	for i := 2; i < n+1; i++ {
		tmp := 0
		// 当前不为 0
		if s[i-1] != '0' {
			tmp += dp2
		}
		if s[i-2] != '0' && ((s[i - 2] - '0') * 10 + (s[i - 1] - '0') <= 26){
			tmp += dp1
		}
		dp1 = dp2
		dp2 = tmp
	}
	return dp2
}
