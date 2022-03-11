package letcode

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var s []byte
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	divide := func(numerator int, denominator int) (a, b int) {
		a = numerator / denominator
		b = numerator % denominator
		return
	}

	// 去掉整数部分
	value, numerator := divide(numerator, denominator)
	s = append(s, strconv.Itoa(value)...)
	if numerator == 0 {
		return string(s)
	}
	s = append(s, '.')

	// 记录上一次出现这个 被除数 的 长度
	set := make(map[int]int)
	for numerator != 0 && set[numerator] == 0 {
		set[numerator] = len(s)
		value, numerator = divide(numerator * 10, denominator)
		// 继续除
		s = append(s, strconv.Itoa(value)...)
	}
	if numerator != 0 {
		// 发现循环
		insertIndex := set[numerator]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
		return string(s)
	}

	return string(s)
}
