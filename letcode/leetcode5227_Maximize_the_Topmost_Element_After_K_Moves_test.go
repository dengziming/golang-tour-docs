package letcode

import (
	"testing"
)

func TestMaximumTop(t *testing.T) {
	if maximumTop([]int{99,95,68,24,18}, 69) != 5 {
		t.Error("测试失败")
	}
	if maximumTop([]int{2}, 1) != -1 {
		t.Error("测试失败")
	}
	if maximumTop([]int{5,2,2,4,0,6}, 4) != 5 {
		t.Error("测试失败")
	}
}
