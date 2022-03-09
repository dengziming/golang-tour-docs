package letcode

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	if len(subsetsWithDup([]int{1, 2, 2})) != 6 {
		fmt.Printf("%v", subsetsWithDup([]int{1, 2, 2}))
		t.Error("测试失败")
	}
}
