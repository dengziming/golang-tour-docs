package letcode

import "testing"

func TestFindKDistantIndices(t *testing.T) {
	if len(findKDistantIndices([]int{2,2,2,2,2}, 2, 2)) != 5 {
		t.Error("测试失败")
	}
	if len(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)) != 6 {
		t.Error("测试失败")
	}

}
