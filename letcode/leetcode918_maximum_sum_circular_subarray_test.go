package letcode

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	if maxSubarraySumCircular([]int{1,-2,3,-2}) != 3 {
		t.Error("测试失败")
	}

	if maxSubarraySumCircular([]int{5,-3,5}) != 10 {
		t.Error("测试失败")
	}

	if maxSubarraySumCircular([]int{-3,-2,-3}) != -2 {
		t.Error("测试失败")
	}
}
