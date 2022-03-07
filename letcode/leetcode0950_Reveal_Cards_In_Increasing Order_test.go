package letcode

import (
	"testing"
)

func TestDeckRevealedIncreasing(t *testing.T) {
	if len(deckRevealedIncreasing([]int{17,13,11,2,3,5,7})) != 7 {
		t.Error("测试失败")
	}
}
