package letcode

import "testing"

func TestIsValid(t *testing.T) {

	if isValid("((") {
		t.Error("测试失败")
	}

}