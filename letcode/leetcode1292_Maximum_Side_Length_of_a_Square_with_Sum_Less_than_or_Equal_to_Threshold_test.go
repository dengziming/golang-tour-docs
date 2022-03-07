package letcode

import "testing"

func TestMaxSideLength(t *testing.T) {
	if maxSideLength([][]int{{1,1,3,2,4,3,2},{1,1,3,2,4,3,2},{1,1,3,2,4,3,2}}, 4) != 2 {
		t.Error("测试失败")
	}
	if maxSideLength([][]int{{2,2,2,2},{2,2,2,2},{2,2,2,2},{2,2,2,2}}, 1) != 0 {
		t.Error("测试失败")
	}
	if maxSideLength([][]int{{18,70},{61,1},{25,85},{14,40},{11,96},{97,96},{63,45}}, 40184) != 2 {
		t.Error("测试失败")
	}

}
