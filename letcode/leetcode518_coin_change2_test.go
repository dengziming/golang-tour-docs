package letcode

import "testing"

func TestChange(t *testing.T) {

	if change(5, []int{1, 2, 5}) != 4 {
		t.Error("测试失败")
	}

	if change(3, []int{2}) != 0 {
		t.Error("测试失败")
	}

	if change(10, []int{10}) != 1 {
		t.Error("测试失败")
	}

}
