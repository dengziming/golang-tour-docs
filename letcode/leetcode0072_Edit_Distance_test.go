package letcode

import "testing"

func TestMinDistance(t *testing.T) {
	if minDistance("", "") != 0 {
		t.Error("测试失败")
	}
	if minDistance("horse", "ros") != 3 {
		t.Error("测试失败")
	}
	if minDistance("intention", "execution") != 5 {
		t.Error("测试失败")
	}
}