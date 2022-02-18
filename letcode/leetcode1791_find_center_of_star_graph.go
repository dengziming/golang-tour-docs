package letcode

func findCenter(edges [][]int) int {
	star1 := edges[0][0]
	star2 := edges[0][1]

	if edges[1][0] == star1 || edges[1][1] == star1 {
		return star1
	}
	return star2
}
