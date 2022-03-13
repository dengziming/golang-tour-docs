package letcode

func minFlips(target string) int {
	flips := 0
	prev := '0'
	for _, curr := range target {
		if curr != prev {
			flips++
			prev = curr
		}
	}

	return flips
}
