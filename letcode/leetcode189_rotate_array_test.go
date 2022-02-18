package letcode

import "testing"

func TestRotate(t *testing.T) {

	arr := []int{1,2,3,4,5,6,7}
	rotate(arr, 3)

	if len(arr) != 7 {
		t.Error("测试失败")
	}
}
