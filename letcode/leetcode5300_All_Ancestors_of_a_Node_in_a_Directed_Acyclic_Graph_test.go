package letcode

import (
	"fmt"
	"testing"
)

func TestGetAncestors(t *testing.T) {

	if len(getAncestors(8, [][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}})) != 8 {
		t.Error("测试失败")
	}
	fmt.Printf("%v", getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))
}
