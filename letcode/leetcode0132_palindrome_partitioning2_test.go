package letcode

import "testing"


func TestMinCut(t *testing.T) {
	if minCut("aab") != 1 {
		t.Error("测试失败")
	}

}
