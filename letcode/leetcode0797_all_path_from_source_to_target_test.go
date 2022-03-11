package letcode

import "testing"

func TestAllPathsSourceTarget(t *testing.T) {
	if len(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}})) != 2 {
		t.Error("测试失败")
	}

	if len(allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{3},{4},{}})) != 5 {
		t.Error("测试失败")
	}
}
