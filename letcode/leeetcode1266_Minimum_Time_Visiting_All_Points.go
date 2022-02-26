package letcode

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	distance := func(a []int, b []int) int {
		x := math.Abs(float64(a[0] - b[0]))
		y := math.Abs(float64(a[1] - b[1]))

		return int(math.Max(x, y))
	}

	d := 0
	for i := 1; i < len(points); i++ {
		d +=distance(points[i], points[i-1])
	}
	return d
}
