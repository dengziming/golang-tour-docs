package letcode

import "testing"

func TestNumTeams(t *testing.T) {
	if numTeams([]int{2,5,3,4,1}) != 3 {
		t.Error("测试失败")
	}
}
