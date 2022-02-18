package letcode

import "testing"

func TestFlipAndInvertImage(t *testing.T) {
	arr := [][]int{{1,1,0},{1,0,1},{0,0,0}}
	/*flipAndInvertImage(arr)
	if len(arr) != 3 {
		t.Error("测试失败")
	}*/

	arr = [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}
	flipAndInvertImage(arr)
	if len(arr) != 4 {
		t.Error("测试失败")
	}
}
