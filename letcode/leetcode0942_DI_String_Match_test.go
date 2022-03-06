package letcode

import "testing"

func TestDiStringMatch(t *testing.T) {

	if len(diStringMatch("IDID")) != 5 {
		t.Error("测试失败")
	}

}
