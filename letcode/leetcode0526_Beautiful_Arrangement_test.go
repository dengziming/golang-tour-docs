package letcode

import (
	"fmt"
	"testing"
)

func TestCountArrangement(t *testing.T) {
	fun := func(x int) int {
		return x & -x
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d, %d\n", i, fun(i))
	}
}
