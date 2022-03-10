package letcode

import "testing"

func TestIncreasingTripletg(t *testing.T) {
	if !increasingTriplet([]int{1, 2, 3}) {
		t.Error("测试失败")
	}

	if increasingTriplet([]int{5,4,3,2,1}) {
		t.Error("测试失败")
	}

	if !increasingTriplet([]int{1,2,3,4,5}) {
		t.Error("测试失败")
	}

	if increasingTriplet([]int{0,4,2,1,0,-1,-3}) {
		t.Error("测试失败")
	}

}
