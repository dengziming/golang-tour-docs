package letcode

import "testing"

func TestValidUtf8(t *testing.T) {

	if !validUtf8([]int{197,130,1}){
		t.Error("测试失败")
	}
	if validUtf8([]int{235,140,4}){
		t.Error("测试失败")
	}
}
