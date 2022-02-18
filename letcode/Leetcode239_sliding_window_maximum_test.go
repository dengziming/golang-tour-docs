package letcode

import "testing"


func TestMaxSlidingWindow(t *testing.T) {

	if len(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3)) != 6 {
		// 3,3,5,5,6,7
		t.Error("测试失败")
	}

	if len(maxSlidingWindow([]int{1}, 1)) != 1 {
		// 1
		t.Error("测试失败")
	}

}
