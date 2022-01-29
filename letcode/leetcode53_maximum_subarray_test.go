package letcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	if maxSubArray([]int{1, 2, 3}) != 6 {
		t.Error("测试失败")
	}

	if maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) != 6 {
		t.Error("测试失败")
	}
	if maxSubArray([]int{1}) != 1 {
		t.Error("测试失败")
	}
	if maxSubArray([]int{5,4,-1,7,8}) != 23 {
		t.Error("测试失败")
	}
}
