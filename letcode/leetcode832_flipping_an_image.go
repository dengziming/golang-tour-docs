package letcode


func flipAndInvertImage(image [][]int) [][]int {

	flip := func(a int) int{
		if a == 0 {
			return 1
		}
		return 0
	}

	revereFlip := func(row *[]int, a int, b int) {
		tmp := (*row)[a]
		(*row)[a] = flip((*row)[b])
		(*row)[b] = flip(tmp)
	}

	n := len(image)
	for _, row := range image {

		for i := 0; i < (n + 1) / 2; i++ {
			revereFlip(&row, i, n - 1 - i)
		}
	}
	return image
}
