package letcode

func numSpecial(mat [][]int) int {

	checkRow := func(i int) int {
		flag := false
		tmp := -1
		for idx, num := range mat[i] {
			if !flag && num != 0 {
				flag = true
				tmp = idx
			} else if flag && num != 0 {
				return -1
			}
		}
		if flag {
			return tmp
		}
		return -1
	}

	// 这一列 除了 i 是否全为 0
	checkColumn := func(i int, j int) bool {
		for k := 0; k < len(mat); k++ {
			if k == i {
				continue
			} else if mat[k][j] != 0 {
				return false
			}
		}
		return true
	}

	result := 0
	for i := range mat {
		idx := checkRow(i)
		if idx != -1 {
			if checkColumn(i, idx) {
				result ++
			}
		}
	}
	return result
}
