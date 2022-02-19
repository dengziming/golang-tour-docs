package letcode

import "testing"

func TestFindEvenNumbers(t *testing.T) {

	if len(findEvenNumbers([]int{2,1,3,0})) != 10 {
		t.Error("测试失败")
	}

}
