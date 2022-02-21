package letcode

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	if scoreOfParentheses("(())") != 2 {
		t.Error("测试失败")
	}

}
