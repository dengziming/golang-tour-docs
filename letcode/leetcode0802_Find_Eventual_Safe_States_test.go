package letcode

import (
	"fmt"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	arr := [][]int{{1,2},{2,3},{5},{0},{5},{},{}}

	if len(eventualSafeNodes(arr)) != 4 {
		fmt.Printf("%v", eventualSafeNodes(arr))
		t.Error("测试失败")
	}

	arr = [][]int{{1,2,3,4},{1,2},{3,4},{0,4},{}}
	if len(eventualSafeNodes(arr)) != 1 {
		t.Error("测试失败")
	}

	if len(eventualSafeNodes([][]int{{}, {0,2,3,4}, {3}, {4}, {}})) != 5 {
		t.Error("测试失败")
	}
}
