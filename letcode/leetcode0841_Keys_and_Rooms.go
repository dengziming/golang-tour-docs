package letcode


func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	flag := make([]int, n)
	var dfs func(i int)
	dfs = func(i int) {
		if flag[i] == 0 {
			flag[i] = 1
			for _, room := range rooms[i] {
				dfs(room)
			}
		}
	}

	dfs(0)

	for _, f := range flag {
		if f == 0 {
			return false
		}
	}
	return true
}
