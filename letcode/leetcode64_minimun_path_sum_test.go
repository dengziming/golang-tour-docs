package letcode

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	v1 := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	if v1 != 7 {
		println(v1)
		t.Error("测试失败")
	}

	v2 := minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})
	if v2 != 12 {
		t.Error("测试失败")
	}

	if minPathSum([][]int{{1, 2, 5}, {3, 2, 1}}) != 6 {
		t.Error("测试失败")
	}

	if minPathSum([][]int{{0}}) != 0 {
		t.Error("测试失败")
	}

}
