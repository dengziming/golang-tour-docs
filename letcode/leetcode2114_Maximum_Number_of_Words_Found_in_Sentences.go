package letcode

import "strings"

func mostWordsFound(sentences []string) int {
	maxV := 0
	current := 0

	max := func(a int, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	for _, sentence := range sentences {
		current = len(strings.Split(sentence, " "))
		maxV = max(maxV, current)
	}
	return maxV
}
