package letcode

import "testing"

func TestNumberOfGoodSubsets(t *testing.T) {
	if numberOfGoodSubsets([]int{1,2,3,4}) != 6 {
		t.Error("测试失败")
	}

	if numberOfGoodSubsets([]int{4,2,3,15}) != 5 {
		t.Error("测试失败")
	}

	if numberOfGoodSubsets([]int{5,10,1,26,24,21,24,23,11,12,27,4,17,16,2,6,1,1,6,8,13,30,24,20,2,19}) != 5368 {
		println(numberOfGoodSubsets([]int{5,10,1,26,24,21,24,23,11,12,27,4,17,16,2,6,1,1,6,8,13,30,24,20,2,19}))
		t.Error("测试失败")
	}

	if numberOfGoodSubsets([]int{9,3,14,12,14,3,23,23,30,9,2,6,26,17,5,8,23,6,8,9,2,4,30,21,19,8,1,23,22,26,17,20,5,15,18,20,22,2,15,8,21,9,20}) != 9958 {
		println(numberOfGoodSubsets([]int{9,3,14,12,14,3,23,23,30,9,2,6,26,17,5,8,23,6,8,9,2,4,30,21,19,8,1,23,22,26,17,20,5,15,18,20,22,2,15,8,21,9,20}))
		t.Error("测试失败")
	}

	if numberOfGoodSubsets([]int{10,11,5,1,10,1,3,1,26,11,6,1,1,15,1,7,22,1,1,1,1,1,23,1,29,5,6,1,1,29,1,1,21,19,1,1,1,2,1,11,1,15,1,22,14,1,1,1,1,6,7,1,14,3,5,1,22,1,1,1,17,1,29,2,1,15,10,1,5,7,1,1,1,30,1,30,1,21,10,1,1,1,1,1,2,6,5,7,3,1,1,19,29,1,7,13,14,1,5,26,19,11,1,1,1,1,1,1,1,1,22,15,1,1,13,1,17,1,1,1,13,6,1,10,1,1,17,1,1,3,14,7,17,1,13,1,1,1,1,1,11,1,1,6,1,1,1,1,1,2,1,30,2,26,1,1,14,1,26,29,30,1,13,21,1,1,14,21,1,23,1,15,23,21,1,30,19,19,1,10,23,3,3,17,22,2,26,1,11,1,23,1,1,1,15,1,1,13,1,1}) != 520317213 {
		println(numberOfGoodSubsets([]int{10,11,5,1,10,1,3,1,26,11,6,1,1,15,1,7,22,1,1,1,1,1,23,1,29,5,6,1,1,29,1,1,21,19,1,1,1,2,1,11,1,15,1,22,14,1,1,1,1,6,7,1,14,3,5,1,22,1,1,1,17,1,29,2,1,15,10,1,5,7,1,1,1,30,1,30,1,21,10,1,1,1,1,1,2,6,5,7,3,1,1,19,29,1,7,13,14,1,5,26,19,11,1,1,1,1,1,1,1,1,22,15,1,1,13,1,17,1,1,1,13,6,1,10,1,1,17,1,1,3,14,7,17,1,13,1,1,1,1,1,11,1,1,6,1,1,1,1,1,2,1,30,2,26,1,1,14,1,26,29,30,1,13,21,1,1,14,21,1,23,1,15,23,21,1,30,19,19,1,10,23,3,3,17,22,2,26,1,11,1,23,1,1,1,15,1,1,13,1,1}))
		t.Error("测试失败")
	}

}
