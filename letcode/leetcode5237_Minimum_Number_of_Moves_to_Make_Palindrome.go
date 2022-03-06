package letcode

import "fmt"

func minMovesToMakePalindrome(s string) int {
	// 贪心，从最外层开始玩
	if len(s) <= 2 {
		return 0
	}
	start := s[0]
	end := s[len(s)-1]
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == end {
			fmt.Printf("%d %s\n", i, s)
			return i + minMovesToMakePalindrome(s[0:i] + s[i+1:len(s)-1])
		} else if start == s[len(s)-i-1] {
			fmt.Printf("%d %s\n", i, s)
			return i + minMovesToMakePalindrome(s[1:len(s)-i-1] + s[len(s)-i:])
		}
	}

	return 0
}
