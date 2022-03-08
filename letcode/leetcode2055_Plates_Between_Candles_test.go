package letcode

import "testing"

func TestPlatesBetweenCandles(t *testing.T) {

	if len(platesBetweenCandles("**", [][]int{{2,2}})) != 1 {
		t.Error("测试失败")
	}

}
