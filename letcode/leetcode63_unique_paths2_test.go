package letcode

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	v1 := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	if v1 != 2 {
		t.Error("测试失败")
	}

	v2 := uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}})
	if v2 != 1 {
		t.Error("测试失败")
	}

}
