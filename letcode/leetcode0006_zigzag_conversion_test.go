package letcode

import "testing"

func TestConvert(t *testing.T) {
	if convert("PAYPALISHIRING", 3) != "PAHNAPLSIIGYIR" {
		t.Error("测试失败")
	}

	if convert("PAYPALISHIRING", 4) != "PINALSIGYAHRPI" {
		t.Error("测试失败")
	}

	if convert("A", 1) != "A" {
		t.Error("测试失败")
	}

	if convert("AB", 1) != "AB" {
		t.Error("测试失败")
	}
}
