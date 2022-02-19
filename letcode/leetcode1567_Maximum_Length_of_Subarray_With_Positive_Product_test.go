package letcode

import "testing"

func TestGetMaxLen(t *testing.T) {
	if getMaxLen([]int{1,-2,-3,4}) != 4 {
		t.Error("测试失败")
	}

	if getMaxLen([]int{0,1,-2,-3,-4}) != 3 {
		t.Error("测试失败")
	}

	if getMaxLen([]int{-1,-2,-3,0,1}) != 2 {
		t.Error("测试失败")
	}

	if getMaxLen([]int{-1,2}) != 1 {
		t.Error("测试失败")
	}
}

