package letcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	if findMedianSortedArrays([]int{1, 2}, []int{3}) != 2 {
		t.Error("测试失败")
	}

	if findMedianSortedArrays([]int{1, 2}, []int{3, 4}) != 2.5 {
		t.Error("测试失败")
	}
}