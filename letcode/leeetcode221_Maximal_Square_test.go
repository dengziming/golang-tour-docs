package letcode

import "testing"

func TestMaximalSquare(t *testing.T) {
	/*if maximalSquare([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}) != 4 {
		t.Error("测试失败")
	}
	if maximalSquare([][]byte{{'0'},{'1'}}) != 1 {
		t.Error("测试失败")
	}*/
	if maximalSquare([][]byte{{'0','0','0','1'},{'1','1','0','1'},{'1','1','1','1'},{'0','1','1','1'},{'0','1','1','1'}}) != 9 {
		t.Error("测试失败")
	}
}
