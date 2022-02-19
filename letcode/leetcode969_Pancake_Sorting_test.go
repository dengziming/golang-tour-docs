package letcode

import "testing"


func TestPancakeSort(t *testing.T) {
	if len(pancakeSort([]int{3,2,4,1})) != 6 {
		t.Error("测试失败")
	}
	if len(pancakeSort([]int{2,1,3})) != 6 {
		t.Error("测试失败")
	}
}
