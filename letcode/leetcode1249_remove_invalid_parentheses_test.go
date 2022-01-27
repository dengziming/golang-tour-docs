package letcode

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	if minRemoveToMakeValid("a()b()c()d)") != "a()b()c()d" {
		t.Error("测试失败")
	}

	if minRemoveToMakeValid("a()b()c()d(") != "a()b()c()d" {
		t.Error("测试失败")
	}

	if minRemoveToMakeValid("a()b()c()d(e") != "a()b()c()de" {
		t.Error("测试失败")
	}
}