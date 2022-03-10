package letcode

// https://leetcode-cn.com/problems/integer-break/
func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	pow := func(a, b int) int {
		ans := 1
		for b > 0 {
		if b & 1 == 1 {
		ans *= a
	}
		a *= a
		b >>= 1
	}
		return ans
	}

	if r := n % 3; r == 0 {
		return pow(3, n / 3)
	} else if r == 1 {
		return pow(3, (n - 4) / 3) * 4
	} else {
		return pow(3, (n - 2) / 3) * 2
	}
}
