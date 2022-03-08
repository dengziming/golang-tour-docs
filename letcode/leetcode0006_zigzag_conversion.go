package letcode

func convert(s string, numRows int) string {
	result := make([]string, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = ""
	}

	x := 0
	direct := 1 // 代表相下
	for _, char:= range s {
		result[x] = result[x]+string(char)

		// 是否要改变 x 的值
		if numRows == 1 {
			direct = 0
		} else if x == 0 {
			direct = 1
		} else if x == numRows - 1 {
			direct = -1
		}
		x = x + direct
	}
	ans := ""
	for _, str := range result {
		ans = ans+str
	}
	return ans
}
