package letcode

import "testing"

func TestMaximumRequests(t *testing.T) {
	if maximumRequests(5, [][]int{{0,1},{1, 0}, {0,1},{1,2},{2,0}, {3,4}}) != 5 {
		t.Error("测试失败")
	}
	if maximumRequests(5, [][]int{{1,2},{1, 2}, {2,2},{0,2},{2,1}, {1,1}, {1,2}}) != 4 {
		t.Error("测试失败")
	}
	if maximumRequests(5, [][]int{{0,1},{1,0}, {0,1},{1,2},{2,0}, {3,4}}) != 5 {
		t.Error("测试失败")
	}
}
