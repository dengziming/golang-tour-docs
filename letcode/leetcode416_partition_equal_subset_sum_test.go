package letcode

import "testing"

func TestCanPartition(t *testing.T) {
	if !canPartition([]int{1, 2, 3}){
		t.Error("测试失败")
	}

	if canPartition([]int{1, 2, 5}){
		t.Error("测试失败")
	}

	if !canPartition([]int{3,3,3,4,5}){
		t.Error("测试失败")
	}

}
