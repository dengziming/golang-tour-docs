package letcode

import "testing"

func TestGenerateParenthesis2(t *testing.T) {
	if len(generateParenthesis2(3)) != 5 {
		t.Error("测试失败")
	}

	if len(generateParenthesis3(3)) != 5 {
		t.Error("测试失败")
	}

}