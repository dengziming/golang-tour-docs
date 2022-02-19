package letcode

import "testing"


func TestMakeConnected(t *testing.T) {
	if makeConnected(4, [][]int{{0,1}, {0,2}, {1,2}}) != 1 {
		t.Error("测试失败")
	}

	if makeConnected(6, [][]int{{0,1}, {0,2}, {0,3}, {1,2}}) != -1 {
		t.Error("测试失败")
	}
}
