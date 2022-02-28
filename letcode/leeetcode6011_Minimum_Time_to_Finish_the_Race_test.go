package letcode

import (
	"testing"
)

func TestMinimumFinishTime(t *testing.T) {
	if minimumFinishTime([][]int{{2,3},{3,4}}, 5, 4) != 21 {
		t.Error("测试失败")
	}
	if minimumFinishTime([][]int{{99,7}}, 85, 95) != 21 {
		t.Error("测试失败")
	}
	if minimumFinishTime([][]int{{1,2}}, 1, 1) != 1 {
		t.Error("测试失败")
	}
}
