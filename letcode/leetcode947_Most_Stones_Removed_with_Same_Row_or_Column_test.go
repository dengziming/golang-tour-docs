package letcode

import "testing"

func TestRemoveStones(t *testing.T) {
	if removeStones([][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}}) != 5 {
		t.Error("测试失败")
	}

}
