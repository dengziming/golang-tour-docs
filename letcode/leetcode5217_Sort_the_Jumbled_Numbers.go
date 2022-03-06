package letcode

import (
	"sort"
	"strconv"
)

func sortJumbled(mapping []int, nums []int) []int {
	n := len(nums)
	arr := make([][]int, n)
	var transform func(x int) int
	transform = func(x int) int {
		xStr := strconv.Itoa(x)
		var xt int
		for _, c := range xStr {
			xt *= 10
			xt += mapping[c-'0']
		}
		return xt
	}
	for i := 0; i < n; i++ {
		arr[i] = []int{nums[i], transform(nums[i])}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = arr[i][0]
	}
	return res
}
