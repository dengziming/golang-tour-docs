package letcode

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	if longestValidParentheses("") + longestValidParentheses("(") + longestValidParentheses(")") != 0 {
		t.Error("测试失败")
	}

	if longestValidParentheses("()") != 2 {
		t.Error("测试失败")
	}

	if longestValidParentheses(")()())") != 4 {
		t.Error("测试失败")
	}

	if longestValidParentheses("(()") != 2 {
		t.Error("测试失败")
	}

	if longestValidParentheses("())") != 2 {
		t.Error("测试失败")
	}

	if longestValidParentheses("()(())") != 6 {
		t.Error("测试失败")
	}

	if longestValidParentheses("(())") != 4 {
		t.Error("测试失败")
	}
}
