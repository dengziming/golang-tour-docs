package letcode

import "testing"


func TestCanConstruct(t *testing.T) {
	if !canConstruct("aa", "aab") {
		t.Error("测试失败")
	}

	if canConstruct("a", "b") {
		t.Error("测试失败")
	}

	if canConstruct("aa", "ab") {
		t.Error("测试失败")
	}
}
