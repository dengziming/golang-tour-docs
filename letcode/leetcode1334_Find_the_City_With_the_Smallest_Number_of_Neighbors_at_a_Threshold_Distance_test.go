package letcode

import (
	"testing"
)

func TestFindTheCity(t *testing.T) {
	if findTheCity(4, [][]int{{0,1,3}, {1,2,1}, {1,3,4}, {2,3,1}}, 4) != 3 {
		t.Error("测试失败")
	}

	if findTheCity(5, [][]int{{0,1,2}, {0,4,8}, {1,2,3}, {1,4,2}, {2,3,1}, {3,4,1}}, 2) != 0 {
		t.Error("测试失败")
	}
}

