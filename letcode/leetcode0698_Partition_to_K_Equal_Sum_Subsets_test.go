package letcode

import (
	"testing"
)

func TestCanPartitionKSubsets(t *testing.T) {
	if !canPartitionKSubsets([]int{4,5,9,3,10,2,10,7,10,8,5,9,4,6,4,9}, 5) {
		t.Error("测试失败")
	}
	if !canPartitionKSubsets([]int{4,4,6,2,3,8,10,2,10,7}, 4) {
		t.Error("测试失败")
	}
	if !canPartitionKSubsets([]int{4,3,2,3,5,2,1}, 4) {
		t.Error("测试失败")
	}
}
