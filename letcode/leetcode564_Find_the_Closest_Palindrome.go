package letcode

import (
	"math"
	"strconv"
	"strings"
)

func nearestPalindromic(n string) string {
	if len(n) == 1 {
		i := n[0] - '0'
		if i == 0 {
			return "1"
		}
		return strconv.Itoa(int((i - 1) % 10))
	}
	getNum := func(prefix string, odd bool) int {
		b := strings.Builder{}
		b.WriteString(prefix)
		length := len(prefix)
		idx := length - 1
		if odd {
			idx --
		}
		for ;idx >= 0; idx-- {
			b.WriteByte(prefix[idx])
		}
		i,_ := strconv.Atoi(b.String())
		return i
	}

	s := make([]byte, len(n))

	result := make([]int, 0)
	// 对于 99，要考虑 101
	result = append(result, int(math.Pow10(len(s))) + 1)
	// 对于 101，要考虑 99
	result = append(result, int(math.Pow10(len(s)-1)) - 1)

	prefix,_ := strconv.Atoi(n[0:(len(n)+1)/2])

	for i := -1; i <= 1; i++ {
		result = append(result, getNum(strconv.Itoa(prefix+i), len(n) % 2 == 1))
	}

	// 选择最小
	min := result[0]
	v,_ := strconv.Atoi(n)
	for i := 1; i < len(result); i++ {
		if math.Abs(float64(result[i]-v)) < math.Abs(float64(min-v)) && result[i] != v {
			min = result[i]
		} else if math.Abs(float64(result[i]-v)) == math.Abs(float64(min-v)) && result[i] < min && result[i] != v {
			min = result[i]
		}
	}
	return strconv.Itoa(min)
}
