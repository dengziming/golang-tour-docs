package letcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abccdd") != 3 {
		t.Error("测试失败")
	}
	if lengthOfLongestSubstring("abcdef") != 6 {
		t.Error("测试失败")
	}
	if lengthOfLongestSubstring("aabcddef") != 4 {
		t.Error("测试失败")
	}
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("测试失败")
	}

}
