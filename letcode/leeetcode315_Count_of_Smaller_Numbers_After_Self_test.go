package letcode

import "testing"

func TestCountSmaller(t *testing.T) {
	if len(countSmaller([]int{5, 2, 6, 1})) != 4 {
		t.Error("测试失败")
	}

}
