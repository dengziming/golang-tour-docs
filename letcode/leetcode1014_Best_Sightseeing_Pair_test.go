package letcode

import "testing"


func TestMaxScoreSightseeingPair(t *testing.T) {
	if maxScoreSightseeingPair([]int{8,1,5,2,6}) != 11 {
		t.Error("测试失败")
	}

	if maxScoreSightseeingPair([]int{1,2}) != 2 {
		t.Error("测试失败")
	}
}
