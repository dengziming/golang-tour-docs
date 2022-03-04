package letcode

func numTilePossibilities(tiles string) int {
	counts := make(map[byte]int)
	chars := make(map[byte]bool)

	for i := range tiles {
		counts[tiles[i]]++
		chars[tiles[i]] = true
	}

	var dfs func(i int)

	cnt := 0
	dfs = func(i int) {
		if i == len(tiles) {
			return
		}

		for b := range chars {
			if counts[b] > 0 {
				counts[b]--
				cnt++
				dfs(i+1)
				counts[b]++
			}
		}
	}
	dfs(0)
	return cnt
}
