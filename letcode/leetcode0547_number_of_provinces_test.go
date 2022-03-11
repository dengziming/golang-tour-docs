package letcode

import "testing"


func TestFindCircleNum(t *testing.T) {
	if findCircleNum([][]int{{1,1,0}, {1,1,0}, {0,0,1}}) != 2 {
		t.Error("测试失败")
	}

	if findCircleNum([][]int{{1,0,0}, {0,1,0}, {0,0,1}}) != 3 {
		t.Error("测试失败")
	}

	if findCircleNum([][]int{{1,0,0,1}, {0,1,1,0}, {0,1,1,1}, {1,0,1,1}}) != 1 {
		t.Error("测试失败")
	}
}
