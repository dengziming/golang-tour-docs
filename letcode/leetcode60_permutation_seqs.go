package letcode

import (
	"fmt"
	"strconv"
	"strings"
)

// 首先需要一个长度为 n 的数组保存 n! 的值
// 对于总的排列，可以分为 n! 组
func getPermutation(n int, k int) string {
	indexes := make([]string, n)

	rest := make([]int, n)
	for i := 0; i <= n - 1; i++ {
		rest[i] = i + 1
	}
	getPermutation1(n, k - 1, 0, &indexes, rest)
	fmt.Printf("%v\n", indexes)

	return strings.Join(indexes, "")
}

// n 总数
// k 求的第几个，从 0 开始编号
// 当前得到的编号
func getPermutation1(n int, k int, i int, indexes *[]string, rest []int) *[]string {
	if i == n {
		return indexes
	}

	totalGroup := multiple(n - 1 - i)
	// 得到 k 属于第几组，以及组内编号
	group := k / totalGroup
	num := k % totalGroup

	(*indexes)[i] = strconv.Itoa(rest[group])
	rest = append(rest[0: group], rest[group + 1: n]...)
	return getPermutation1(n, num, i + 1, indexes, rest)
}

func multiple(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i ++ {
		result = result * i
	}
	return result
}
