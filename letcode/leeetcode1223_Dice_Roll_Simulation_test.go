package letcode

import (
	"testing"
)

func TestDieSimulator(t *testing.T) {
	if dieSimulator(30, []int{2,3,1,2,1,2}) != 753152086 {
		println(dieSimulator(30, []int{2,3,1,2,1,2}))
		t.Error("测试失败")
	}
	if dieSimulator(20, []int{8,5,10,8,7,2}) != 822005673 {
		println(dieSimulator(20, []int{8,5,10,8,7,2}))
		t.Error("测试失败")
	}
	if dieSimulator(100, []int{8,11,13,9,10,7}) != 778061023 {
		println(dieSimulator(100, []int{8,11,13,9,10,7}))
		t.Error("测试失败")
	}


}
