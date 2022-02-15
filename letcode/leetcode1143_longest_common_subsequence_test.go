package letcode

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	if longestCommonSubsequence("abcde", "ace") != 3 {
		t.Error("测试失败")
	}

	if longestCommonSubsequence("abc", "abc") != 3 {
		t.Error("测试失败")
	}

	if longestCommonSubsequence("abc", "def") != 0 {
		t.Error("测试失败")
	}
}
