package letcode

import (
	"testing"
)

func TestMinimalKSum(t *testing.T) {
	if minimalKSum([]int{1,4,25,10,25}, 2) != 5 {
		t.Error("测试失败")
	}

}
