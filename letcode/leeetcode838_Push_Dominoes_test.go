package letcode

import "testing"

func TestPushDominoes(t *testing.T) {
	if pushDominoes(".") != "." {
		println(pushDominoes("RR.L"))
		t.Error("测试失败")
	}
	if pushDominoes(".L.R.") != "LL.RR" {
		println(pushDominoes("RR.L"))
		t.Error("测试失败")
	}
	if pushDominoes("RR.L") != "RR.L" {
		println(pushDominoes("RR.L"))
		t.Error("测试失败")
	}
	if pushDominoes(".L.R...LR..L..") != "LL.RR.LLRRLL.." {
		t.Error("测试失败")
	}
}
