package letcode

import "testing"

func TestSubArrayRanges(t *testing.T) {
	if subArrayRanges([]int{4,-2,-3,4,1}) != 59 {
		t.Error("测试失败")
	}
	if subArrayRanges([]int{1,2,3}) != 4 {
		t.Error("测试失败")
	}
	if subArrayRanges([]int{1,3,3}) != 4 {
		t.Error("测试失败")
	}
}
