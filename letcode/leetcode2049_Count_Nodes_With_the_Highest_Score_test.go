package letcode

import "testing"


func TestCountHighestScoreNodes(t *testing.T) {
	if countHighestScoreNodes([]int{-1,2,0}) != 2 {
		t.Error("测试失败")
	}
	if countHighestScoreNodes([]int{-1,2,0,2,0}) != 3 {
		t.Error("测试失败")
	}
}
