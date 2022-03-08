package letcode

import "testing"

func TestCheckArithmeticSubarrays(t *testing.T) {
	if !checkArithmeticSubarrays([]int{-3,-6,-8,-4,-2,-8,-6,0,0,0,0}, []int{7}, []int{10})[0] {
		t.Error("测试失败")
	}
	if len(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5})) != 3 {
		t.Error("测试失败")
	}

}
