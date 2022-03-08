package letcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	abs := func(a int, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	sort.Ints(seats)
	sort.Ints(students)
	result := 0
	for i := 0; i < len(students); i++ {
		result += abs(students[i], seats[i])
	}
	return result
}
