package letcode

import "testing"

func TestThreeSum(t *testing.T) {
	if len(threeSum([]int{2, 3, 6, 7})) != 0 {
		printlnArray(threeSum([]int{2, 3, 6, 7}))
		t.Error("测试失败")
	}

	if len(threeSum([]int{-1,0,1,2,-1,-4})) != 2 {
		printlnArray(threeSum([]int{-1,0,1,2,-1,-4}))
		t.Error("测试失败")
	}
}
