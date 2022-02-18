package letcode

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	if len(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})) != 2 {
		t.Error("测试失败")
	}

	if len(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})) != 2 {
		t.Error("测试失败")
	}
}
