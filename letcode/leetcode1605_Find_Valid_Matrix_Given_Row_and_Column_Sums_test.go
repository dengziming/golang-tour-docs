package letcode

import "testing"

func TestRestoreMatrix(t *testing.T) {
	if len(restoreMatrix([]int{3, 8}, []int{4, 7})) != 2 {
		t.Error("测试失败")
	}
}
