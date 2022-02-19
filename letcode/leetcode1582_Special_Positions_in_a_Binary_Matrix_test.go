package letcode

import "testing"

func TestNumSpecial(t *testing.T) {
	if numSpecial([][]int{{1,0,0},{0,0,1},{1,0,0}}) != 1 {
		t.Error("测试失败")
	}

}
