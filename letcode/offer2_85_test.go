package letcode

import "testing"

func TestGenerateParenthesis2(t *testing.T) {
	if len(generateParenthesis2(3)) != 5 {
		t.Error("测试失败")
	}

	if len(generateParenthesis3(3)) != 5 {
		t.Error("测试失败")
	}

	if maxTrailingZeros([][]int{{1,5,2,4,25}}) != 3 {
		t.Error("测试失败")
	}
	if maxTrailingZeros([][]int{{23,17,15,3,20}, {8,1,20,27,11}, {9,4,6,2,21}, {40,9,1,10,6}, {22,7,4,5,3}}) != 3 {
		t.Error("测试失败")
	}

	if len(countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}})) != 2 {
		t.Error("测试失败")
	}
}