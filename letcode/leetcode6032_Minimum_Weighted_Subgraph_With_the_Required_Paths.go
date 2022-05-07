package letcode

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	weights := make(map[[2]int]int)
	inEdges := make([][]int, n)
	for _, edge := range edges {
		inEdges[edge[1]] = append(inEdges[edge[1]], edge[0])
		weights[[2]int{edge[0], edge[1]}] = edge[2]
	}


	return 0

}
