package letcode

import "math"

func subsets(nums []int) [][]int {
	n := len(nums)
	cnt := int(math.Pow(2, float64(n)))

	result := make([][]int, 0)
	for i := 0; i < cnt; i++ {
		arr := make([]int, 0)
		for j := 0; j < n; j++ {
			// i 的二进制表示的第 j 位是否是 1
			if (i >> j)%2 == 1 {
				arr = append(arr, nums[j])
			}
		}
		result = append(result, arr)
	}
	return result
}
