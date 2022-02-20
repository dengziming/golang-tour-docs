package letcode

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	flag := false
	for i < n {
		char := bits[i]
		if char == 1 {
			i += 2
			flag = false
		} else {
			i += 1
			flag = true
		}
	}

	return flag
}
