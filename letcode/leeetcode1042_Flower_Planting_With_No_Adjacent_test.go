package letcode

import (
	"testing"
)

func TestGardenNoAdj(t *testing.T) {

	if len(gardenNoAdj(4, [][]int{{1,2},{3,4}})) != 4 {
		t.Error("测试失败")
	}

}
