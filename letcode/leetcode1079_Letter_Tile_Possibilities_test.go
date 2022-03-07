package letcode

import (
	"fmt"
	"testing"
)

func TestNumTilePossibilities(t *testing.T) {
	if numTilePossibilities("AAB") != 8 {
		fmt.Printf("%v", numTilePossibilities("AAB"))
		t.Error("测试失败")
	}

}
