package letcode

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	if !canVisitAllRooms([][]int{{1},{2},{3},{}}) {
		t.Error("测试失败")
	}

	if canVisitAllRooms([][]int{{1,3},{3,0,1},{2},{0}}) {
		t.Error("测试失败")
	}

}
