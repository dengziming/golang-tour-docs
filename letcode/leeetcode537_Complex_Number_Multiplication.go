package letcode

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	mul := func(a1, a2, b1, b2 int) (int, int) {
		return a1 * b1 - a2 * b2, a1 * b2 + a2 * b1
	}

	convert := func(num string) (int, int) {
		a := strings.Split(num, "+")
		v1, _ := strconv.Atoi(a[0])
		v2, _ := strconv.Atoi(strings.Split(a[1], "i")[0])
		return v1, v2
	}

	v11, v12 := convert(num1)
	v21, v22 := convert(num2)
	v1, v2 := mul(v11,v12,v21,v22)
	return fmt.Sprintf("%d+%di", v1, v2)

}
