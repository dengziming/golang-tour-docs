package letcode

import (
	"testing"
)


func TestLargestPerimeter(t *testing.T) {
	if largestPerimeter([]int{2,1,2}) != 5 {
		t.Error("测试失败")
	}
	if largestPerimeter([]int{1,2,1}) != 0 {
		t.Error("测试失败")
	}
}
