package letcode

import (
	"strconv"
)

// 思路：卡特兰数，一般两种
func diffWaysToCompute(expression string) []int {
	cache := make(map[string][]int)

	var compute func(expr string) []int
	var op func(left int, right int, o uint8) int
	op = func(left int, right int, o uint8) int {
		if o == '+' {
			return left + right
		}
		if o == '-' {
			return left - right
		}
		if o == '*' {
			return left * right
		}
		return 0
	}

	var isOp func(o uint8) bool
	isOp = func(o uint8) bool {
		return o == '+' || o == '-' || o == '*'
	}

	compute = func(expr string) []int {
		v, isDigit := strconv.Atoi(expr)
		if isDigit == nil {
			return []int{v}
		}
		if _, ok := cache[expr]; ok {
			return cache[expr]
		}
		result := make([]int, 0)
		for i := 0; i < len(expr); i++ {
			if isOp(expr[i]) {
				lefts := compute(expr[0: i])
				rights := compute(expr[i + 1:])
				for _, left := range lefts {
					for _, right := range rights {
						result = append(result, op(left, right, expr[i]))
					}
				}
			}
		}
		cache[expr] = result
		return result
	}
	return compute(expression)
}
