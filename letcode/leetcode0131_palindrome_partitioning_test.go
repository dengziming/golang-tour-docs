package letcode

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	v1 := partition("abb")
	if len(v1) != 2 {
		t.Error("测试失败")
	}

	v2 := partition("a")
	if len(v2) != 1 {
		t.Error("测试失败")
	}

	v3 := partition("abbab")
	// a, b, b, a, b
	// a, b, bab
	// a, bb, a, b
	// abba, b
	if len(v3) != 4 {
		t.Error("测试失败")
	}

	v4 := partition("ababbbabbaba")

	if len(v4) != 94 {
		fmt.Printf("%d, %v\n", len(v4), v4)
		t.Error("测试失败")
	}

}
