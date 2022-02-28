package letcode

import (
	"testing"
)


func TestMinOperations(t *testing.T) {
	/*if minOperations([]int{1,2,3,5,6}) != 1 {
		t.Error("测试失败")
	}
	if minOperations([]int{8,5,9,9,8,4}) != 2 {
		t.Error("测试失败")
	}*/
	if minOperations([]int{8,10,16,18,10,10,16,13,13,16}) != 6 {
		t.Error("测试失败")
	}

}
