package letcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	if longestPalindrome("babad") != "bab" {
		t.Error("测试失败")
	}
	if longestPalindrome("ac") != "a" {
		t.Error("测试失败")
	}
}
