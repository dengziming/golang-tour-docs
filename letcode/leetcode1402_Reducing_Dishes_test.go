package letcode

import (
	"testing"
)

func TestMaxSatisfaction(t *testing.T) {
	if maxSatisfaction([]int{-1,-8,0,5,-7}) != 14 {
		t.Error("测试失败")
	}
	if maxSatisfaction([]int{4,3,2}) != 20 {
		t.Error("测试失败")
	}
	if maxSatisfaction([]int{-1,-4,-5}) != 0 {
		t.Error("测试失败")
	}
}
