package letcode

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	if len(generateParenthesis(1)) != 1 {
		t.Error("测试失败")
	}
	if len(generateParenthesis(2)) != 2 {
		t.Error("测试失败")
	}
	if len(generateParenthesis(3)) != 5 {
		t.Error("测试失败")
	}
}
