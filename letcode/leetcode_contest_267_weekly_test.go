package letcode

import (
	"fmt"
	"testing"
)


func TestMaximumBobPoints(t *testing.T) {
	fmt.Printf("%v", substes([]int{1,2,3}))
}

func substes(nums []int) [][]int {
	subsets := [][]int{}

	n := len(nums)

	cnt := 1
	for i := 0; i < n; i++ {
		cnt = cnt * 2
	}

	for i := 0; i < cnt; i++ {
		result := make([]int, 0)
		for j := 0; j < n; j++ {
			if (i >> j) & 1 == 1 {
				result = append(result, nums[j])
			}
		}
		subsets = append(subsets, result)
	}

	// 
	return subsets
}
