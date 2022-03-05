package letcode

import "testing"

func TestCheckValidString(t *testing.T) {
	if !checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()") {
		t.Error("测试失败")
	}

	if checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())") {
		t.Error("测试失败")
	}

}
