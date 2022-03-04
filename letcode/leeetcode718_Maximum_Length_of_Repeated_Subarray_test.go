package letcode

import "testing"

func TestFindLength(t *testing.T) {
	if findLength([]int{0,1,1,1,1}, []int{1,0,1,0,1}) != 2 {
		t.Error("测试失败")
	}
}
