package letcode

import "testing"


func TestGoodTriplets(t *testing.T) {
	if goodTriplets([]int{2,0,1,3}, []int{0,1,2,3}) != 1 {
		t.Error("测试失败")
	}

	if goodTriplets([]int{4,0,1,3,2}, []int{4,1,0,2,3}) != 4 {
		t.Error("测试失败")
	}
}
