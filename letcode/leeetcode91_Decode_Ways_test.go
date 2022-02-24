package letcode

import "testing"

func TestNumDecodings(t *testing.T) {
	if numDecodings("06") !=0 {
		t.Error("测试失败")
	}
	if numDecodings("0") !=0 {
		t.Error("测试失败")
	}

	if numDecodings("10") !=1 {
		t.Error("测试失败")
	}
}