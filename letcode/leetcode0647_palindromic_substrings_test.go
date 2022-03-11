package letcode

import "testing"

func TestCountSubstrings(t *testing.T) {
	if countSubstrings("abc") != 3 {
		t.Error("测试失败")
	}

	if countSubstrings("aaa") != 6 {
		println(countSubstrings("aaa"))
		t.Error("测试失败")
	}
}
