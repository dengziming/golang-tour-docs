package letcode

import "testing"

func TestMaximumGood(t *testing.T) {
	if maximumGood([][]int{{2,1,2}, {1,2,2}, {2,0,2}}) != 2 {
		t.Error("测试失败")
	}
	if maximumGood([][]int{{2,0,2,2,0}, {2,2,2,1,2}, {2,2,2,1,2}, {1,2,0,2,2}, {1,0,2,1,2}}) != 2 {
		t.Error("测试失败")
	}
}
