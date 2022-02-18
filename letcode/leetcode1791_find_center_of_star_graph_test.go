package letcode

import "testing"

func TestFindCenter(t *testing.T) {
	if findCenter([][]int{{1,2}, {2, 3}, {4,2}}) != 2 {
		t.Error("测试失败")
	}

	if findCenter([][]int{{1,2}, {5, 1}, {1,3}, {1,4}}) != 1 {
		t.Error("测试失败")
	}

}
