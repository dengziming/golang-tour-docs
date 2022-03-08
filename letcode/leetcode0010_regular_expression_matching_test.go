package letcode

import "testing"

func TestRegularExpressionMatching(t *testing.T) {
	if !isMatch("", "") {
		t.Error("测试失败")
	}

	if isMatch("aa", "a") {
		t.Error("测试失败")
	}

	if !isMatch("aa", "a*") {
		t.Error("测试失败")
	}

	if !isMatch("ab", ".*") {
		t.Error("测试失败")
	}

	if !isMatch("aab", "c*a*b*") {
		t.Error("测试失败")
	}

	if isMatch("mississippi", "mis*is*p*.") {
		t.Error("测试失败")
	}

	if isMatch("ab", ".*c") {
		t.Error("测试失败")
	}

	if isMatch("a", "ab*a") {
		t.Error("测试失败")
	}
}
