package letcode

import "testing"

func TestShortestPalindrome(t *testing.T) {
	if shortestPalindrome("aacecaaa") != "aaacecaaa" {
		t.Error("测试失败")
	}

}
