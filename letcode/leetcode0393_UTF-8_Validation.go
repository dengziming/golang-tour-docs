package letcode

func validUtf8(data []int) bool {

	validate := func(data []int) bool {
		mask1 := 1 << 7
		mask2 := 1 << 7 + 1 << 6
		mask3 := 1 << 7 + 1 << 6 + 1 << 5
		mask4 := 1 << 7 + 1 << 6 + 1 << 5 + 1 << 4
		mask5 := 1 << 7 + 1 << 6 + 1 << 5 + 1 << 4 + 1 << 3

		if len(data) == 0 {
			return false
		}
		if len(data) == 1 {
			return (data[0] & mask1) == 0
		}
		if len(data) == 2 {
			return (data[0] & mask3) == mask2 && (mask2 & data[1]) == mask1
		}
		if len(data) == 3 {
			return (data[0] & mask4) == mask3 && (mask2 & data[1]) == mask1 && (mask2 & data[2]) == mask1
		}
		if len(data) == 4 {
			return (data[0] & mask5) == mask4 && (mask2 & data[1]) == mask1 && (mask2 & data[2]) == mask1 && (mask2 & data[3]) == mask1
		}
		return false
	}

	index := 0
	for index < len(data) {
		flag := false
		for j := 1; index+j <= len(data); j++ {
			if validate(data[index : index+j]) {
				index+=j
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
