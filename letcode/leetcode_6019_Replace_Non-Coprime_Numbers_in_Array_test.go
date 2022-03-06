package letcode

import "testing"

func TestReplaceNonCoprimes(t *testing.T) {
	if len(replaceNonCoprimes([]int{287,41,49,287,899,23,23,20677,5,825})) != 3 {
		t.Error("测试失败")
	}
	if len(replaceNonCoprimes([]int{6, 4, 3, 2, 7, 6, 2})) != 3 {
		t.Error("测试失败")
	}
	if len(replaceNonCoprimes([]int{2,2,1,1,3,3,3})) != 4 {
		t.Error("测试失败")
	}

}

