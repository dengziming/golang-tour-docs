package letcode

import (
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	if fractionToDecimal(4, 333) != "0.(012)" {
		println(fractionToDecimal(4, 333))
		t.Error("测试失败")
	}
	if fractionToDecimal(1, 6) != "0.1(6)" {
		println(fractionToDecimal(1, 6))
		t.Error("测试失败")
	}
}
